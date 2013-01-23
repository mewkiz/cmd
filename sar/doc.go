/*

sar uses regexp to search and replace on provided input provided.

Installation:

	$ go get github.com/mewkiz/cmd/sar

Documentation:

Documentation provided by GoDoc.

- sar: http://godoc.org/github.com/mewkiz/cmd/sar

Usage:

	sar [OPTION]... SEARCH REPLACE [FILE]

Flags:

	-i (default=false)
		Edit file in place.

Examples:

1. Search and replace multiple lines.

	$ echo -e "Testing\n1\n2\n3" | sar "1\n2\n3" "`printf "3\n2\n1"`"
	// input:
	//    Testing
	//    1
	//    2
	//    3
	//
	// output:
	//    Testing
	//    3
	//    2
	//    1

2. Use regexp for search and replace.

	$ sar m[a-z]w$ kiz file.txt
	// input (file.txt):
	//    mewmew
	//
	// output:
	//    mewkiz

*/
package documentation
