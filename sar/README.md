# sar

[![GoDoc](http://godoc.org/github.com/mewkiz/cmd/sar?status.svg)](http://godoc.org/github.com/mewkiz/cmd/sar)

The sar tool performs regexp search and replace on the input.

## Installation

```bash
go get -u github.com/mewkiz/cmd/sar
```

## Usage

```
Usage: sar [OPTION]... SEARCH REPLACE [FILE]

Flags:
  -i    Edit file in place.
```

## Examples

1. Search and replace multiple lines.

```bash
echo -e "Testing\n1\n2\n3" | sar "1\n2\n3" "3\n2\n1"
// Input:
// Testing
// 1
// 2
// 3
//
// Output:
// Testing
// 3
// 2
// 1
```

2. Use regexp for search and replace.

```bash
sar "m[a-z]w$" "kiz" file.txt
// Input (file.txt):
// mewmew
//
// Output:
// mewkiz
```

3. Use regexp to edit a file in place.

```bash
sar -i "foo([0-9]+)bar" "num=\$1" foo.txt
// Input (foo.txt):
// foo1234bar
//
// Output (foo.txt):
// num=1234
```
