sar
===

sar is a tool which uses regexp to search and replace input provided from a file
or standard input.

Installation
------------

    $ go get github.com/mewkiz/cmd/sar
    $ go install github.com/mewkiz/cmd/sar

Usage
-----

    $ echo -e "Testing\n1\n2\n3" | sar "1\n2\n3" "`printf "3\n2\n1"`"
    // input:
    // Testing
    // 1
    // 2
    // 3
    //
    // output:
    // Testing
    // 3
    // 2
    // 1

    $ sar mew$ kiz file.txt
    // file content (file.txt):
    // mewmew
    //
    // output:
    // mewkiz
