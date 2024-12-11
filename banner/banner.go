package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.1"

func PrintVersion() {
	fmt.Printf("Current cspfinder version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
                      ____ _             __           
  _____ _____ ____   / __/(_)____   ____/ /___   _____
 / ___// ___// __ \ / /_ / // __ \ / __  // _ \ / ___/
/ /__ (__  )/ /_/ // __// // / / // /_/ //  __// /    
\___//____// .___//_/  /_//_/ /_/ \__,_/ \___//_/     
          /_/`
	fmt.Printf("%s\n%55s\n\n", banner, "Current cspfinder version "+version)
}
