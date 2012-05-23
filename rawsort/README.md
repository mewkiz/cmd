rawsort
=======

rawsort is a tool which writes the sorted concatenation of all provided files to
standard output. If no file is provided it will read from standard input.

Installation
------------

    $ go get github.com/mewkiz/rawsort
    $ go install github.com/mewkiz/rawsort

Usage
-----

    $ rawsort file1.txt file2.txt
    $ rawsort < file1.txt
