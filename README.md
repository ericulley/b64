# b64 Command Line Tool

This is an extremely simple command line tool that decodes or encodes a string to/from base64 encoding.

## Install

```
git clone https://github.com/ericulley/b64.git
cd b64
go build
go install
```

This installs the `b64` binary to `$GOBIN`, which defaults to `$HOME/go/bin` (make sure it's on your `PATH`).

## Usage

```
b64 -e <string>   # encode a string to base64
b64 -d <string>   # decode base64 to a string
```

## Examples

```
$ b64 -e "hello world"
aGVsbG8gd29ybGQ=

$ b64 -d "aGVsbG8gd29ybGQ="
hello world
```
