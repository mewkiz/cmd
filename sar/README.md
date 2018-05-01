# sar

sar uses regexp to search and replace on provided input provided.

## Installation

```bash
go get github.com/mewkiz/cmd/sar
```

## Documentation

Documentation provided by GoDoc.

- [sar](http://godoc.org/github.com/mewkiz/cmd/sar): Use regexp to search and replace on provided input.

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
sar "m[a-z]w$" "kiz"
// Input (file.txt):
// mewmew
//
// Output:
// mewkiz
```
