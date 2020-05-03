/*
The sar tool performs regexp search and replace on the input.

Installation:

	go get -u github.com/mewkiz/cmd/sar

Usage:

	sar [OPTION]... SEARCH REPLACE [FILE]

Flags:

	-fixed-search
	      interpret SEARCH as a fixed string, not a regular expression
	-i    edit file in place
	-unescape-replace
	      unescape REPLACE string
*/
package main
