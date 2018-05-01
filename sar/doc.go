/*
The sar tool performs regexp search and replace on the input.

Installation:

	go get -u github.com/mewkiz/cmd/sar

Usage:

	sar [OPTION]... SEARCH REPLACE [FILE]

Flags:

	-i    Edit file in place.

Examples:

1. Search and replace multiple lines.

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

2. Use regexp for search and replace.

	sar m[a-z]w$ kiz file.txt
	// Input (file.txt):
	// mewmew
	//
	// Output:
	// mewkiz

3. Use regexp to edit a file in place.

	sar -i "foo([0-9]+)bar" "num=\$1" foo.txt
	// Input (foo.txt):
	// foo1234bar
	//
	// Output (foo.txt):
	// num=1234
*/
package main
