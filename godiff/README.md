godiff
======

godiff is a convenient wrapper around gofmt, applying a set of command line
flags and opening the output (if any) in an editor.

*Note*: This command uses hard coded values so alter the source code as you
prefer.

Installation
------------

    $ go get github.com/mewkiz/cmd/godiff

Usage
-----

    $ find . -type f -name '*.go' -print0 | xargs -0 godiff

    $ godiff source.go
