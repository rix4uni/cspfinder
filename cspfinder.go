package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/rix4uni/cspfinder/banner"
)

var concurrent = flag.Int("concurrent", 50, "Number of concurrent requests")
var timeout = flag.Int("timeout", 15, "Timeout for curl requests in seconds")
var silent = flag.Bool("silent", false, "silent mode.")
var versionFlag = flag.Bool("version", false, "Print the version of the tool and exit.")

// getCSP tries fetching the CSP header using https first, then falls back to http if https fails.
func getCSP(input string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	urlToCheck := normalizeURL(input)

	// Acquire a token from the semaphore
	sem <- struct{}{}

	// Try with https first
	if !tryFetchCSP(urlToCheck) {
		// If https fails, fall back to http
		urlToCheck = strings.Replace(urlToCheck, "https://", "http://", 1)
		tryFetchCSP(urlToCheck)
	}

	// Release the token
	<-sem
}

// normalizeURL ensures the URL starts with https:// if no scheme is provided.
func normalizeURL(input string) string {
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		return "https://" + input
	}
	return input
}

// tryFetchCSP runs the curl command and processes the CSP if found.
func tryFetchCSP(url string) bool {
	// Convert timeout to string
	timeoutStr := fmt.Sprintf("%d", *timeout)

	cmd := exec.Command("curl", "--max-time", timeoutStr, "-s", "-I", url)
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.ToLower(line), "content-security-policy:") {
			processCSP(line)
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading output: %v\n", err)
	}

	return false
}

// processCSP processes the CSP header line and extracts domains.
func processCSP(csp string) {
	parts := strings.Fields(csp)
	for _, part := range parts {
		if strings.Contains(part, ".") {
			part = strings.TrimPrefix(part, "wss://")
			part = strings.TrimPrefix(part, "https://")
			part = strings.TrimPrefix(part, "http://")
			part = strings.Split(part, "/")[0]
                        part = strings.Split(part, ";")[0]
			part = strings.TrimSuffix(part, ";")

			fmt.Println(part)
		}
	}
}

func main() {
	// Parse command-line flags
	flag.Parse()

	if *versionFlag {
		banner.PrintBanner()
		banner.PrintVersion()
		return
	}

	if !*silent {
		banner.PrintBanner()
	}

	// Create a semaphore channel to limit concurrency
	sem := make(chan struct{}, *concurrent)

	var wg sync.WaitGroup

	// Read input line-by-line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		wg.Add(1)
		go getCSP(url, &wg, sem)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
