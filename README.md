# scpgo [![Apache Licensed](https://img.shields.io/badge/license-Apache2.0-blue.svg)](https://raw.githubusercontent.com/raravena80/scpgo/master/LICENSE) [![Build Status](https://travis-ci.org/raravena80/scpgo.svg?branch=master)](https://travis-ci.org/raravena80/scpgo) [![Go Report Card](https://goreportcard.com/badge/github.com/raravena80/scpgo)](https://goreportcard.com/report/github.com/raravena80/scpgo)

Go SCP Implementation using the std golang.org/x/crypto/ssh libraries

## Usage

```
This is an SCP implementation in Go.

Usage:
  scpgo <src> host:<dst> [flags]

Flags:
  -c, --checkKnownHosts   Check known hosts
      --config string     config file (default is $HOME/.scpgo.yaml)
  -h, --help              help for scpgo
  -k, --keyFile string    Use this keyfile to authenticate
  -p, --port int          Port number (default 22)
  -q, --quiet             Quiet mode: disables the progress meter as well as warning and diagnostic messages
  -r, --recursive         Recursive copy
  -v, --verbose           Verbose mode - output differs from normal copier
```
