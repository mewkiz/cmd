package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/mewkiz/pkg/imgutil"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: imgcmp FILE0 FILE1")
	fmt.Fprintln(os.Stderr, "Compare two images pixel by pixel.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "An error message will be displayed if they differ.")
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		return
	}
	err := imgcmp(flag.Arg(0), flag.Arg(1))
	if err != nil {
		log.Fatalln(err)
	}
}

// imgcmp compares the two images and returns an error if they differ.
func imgcmp(imgPath0, imgPath1 string) (err error) {
	img0, err := imgutil.ReadFile(imgPath0)
	if err != nil {
		return err
	}
	img1, err := imgutil.ReadFile(imgPath1)
	if err != nil {
		return err
	}
	err = cmp(img0, img1)
	if err != nil {
		return err
	}
	return nil
}

// cmp compares the two images pixel by pixel and returns an error if they differ.
func cmp(img0, img1 image.Image) (err error) {
	rect0 := img0.Bounds()
	rect1 := img1.Bounds()
	size0 := rect0.Size()
	size1 := rect1.Size()
	if size0 != size1 {
		return fmt.Errorf("image sizes differ - img0: %dx%d, img1: %dx%d.", size0.X, size0.Y, size1.X, size1.Y)
	}
	for x := rect0.Min.X; x < rect0.Max.X; x++ {
		for y := rect0.Min.Y; y < rect0.Max.Y; y++ {
			c0 := img0.At(x, y)
			c1 := img1.At(x, y)
			if !colorEqual(c0, c1) {
				return fmt.Errorf("pixel colors differ at x=%d, y=%d; c0=%#v, c1=%#v", x, y, c0, c1)
			}
		}
	}
	return nil
}

// colorEqual reports whether the given colours are equal.
func colorEqual(c0, c1 color.Color) bool {
	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()
	return r0 == r1 && g0 == g1 && b0 == b1 && a0 == a1
}
