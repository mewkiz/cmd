imgcmp
======

imgcmp compares two images and displays an error message if they differ.

Installation
------------

	go get github.com/mewkiz/cmd/imgcmp

Documentation
-------------

Documentation provided by GoDoc.

- [imgcmp][]: Compare two images and displays an error message if they differ.

[imgcmp]: http://godoc.org/github.com/mewkiz/cmd/imgcmp

Usage
-----

	imgcmp FILE0 FILE1

Examples
--------

1. Images with identical pixel content.

		imgcmp img0.png img1.png

2. Images of different sizes.

		imgcmp img0.png img1.png

	Output:
		image sizes differ - img0: 3264x2448, img1: 3648x2736.

3. Images with different pixel content.

		imgcmp img0.png img1.png

	Output:
		pixel colors differ at x=0, y=0.
