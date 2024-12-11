## cspfinder

Discover new target domains using Content Security Policy

## Installation
```
go install github.com/rix4uni/cspfinder@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/cspfinder/releases/download/v0.0.1/cspfinder-linux-amd64-0.0.1.tgz
tar -xvzf cspfinder-linux-amd64-0.0.1.tgz
rm -rf cspfinder-linux-amd64-0.0.1.tgz
mv cspfinder ~/go/bin/cspfinder
```
Or download [binary release](https://github.com/rix4uni/cspfinder/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/cspfinder.git
cd cspfinder; go install
```

## Usage
```
Usage of cspfinder:
  -concurrent int
        Number of concurrent requests (default 50)
  -silent
        silent mode.
  -timeout int
        Timeout for curl requests in seconds (default 15)
  -version
        Print the version of the tool and exit.
```

## Examples
Single Target:
```
▶ echo "https://www.github.com" | cspfinder -silent
```

Multiple Targets:
```
▶ cat targets.txt
google.com
grow.google
https://www.github.com

▶ cat targets.txt | cspfinder -silent
```