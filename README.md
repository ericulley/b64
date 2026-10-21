# b64 Command Line Tool

This is an extremely simple command line tool that decodes or encodes string to/from base64 encoding.

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