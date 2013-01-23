package main

import "flag"
import "fmt"
import "io"
import "log"
import "os"
import "path"
import "regexp"
import "strings"
import "time"

// When flagForce is set to true, force rename the files.
var flagForce bool

// When flagInsane is set to true, skip sanity checks.
var flagInsane bool

func init() {
	flag.BoolVar(&flagForce, "f", false, "Force rename.")
	flag.BoolVar(&flagInsane, "insane", false, "Disable sanity checks.")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: img_date [OPTION]... [FILE]...")
	fmt.Fprintln(os.Stderr, "Rename files based on their embeded date information.")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Flags:")
	fmt.Fprintln(os.Stderr)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, "  Force rename all png files:")
	fmt.Fprintln(os.Stderr, "    img_date -f *.png")
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "  Print rename shell script:")
	fmt.Fprintln(os.Stderr, "    img_date *.jpg")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var renames []*Rename
	for _, filePath := range flag.Args() {
		rename, err := GetRename(filePath)
		if err != nil {
			log.Println(err)
			continue
		}
		renames = append(renames, rename)
	}
	if len(renames) < 1 {
		return
	}
	// Prevent duplicate file names.
	SetSuffixes(renames)
	for _, rename := range renames {
		if !flagInsane {
			// Perform sanity check on date.
			err := rename.SanityCheck()
			if err != nil {
				fmt.Printf("Ignoring %q. Use -insane to disable the sanity check.\n", rename.origPath)
				log.Println(err)
				continue
			}
		}
		if flagForce {
			// Force rename.
			err := os.Rename(rename.origPath, rename.String())
			if err != nil {
				log.Println(err)
			}
		} else {
			// Print a shell script line which renames the original file to the new
			// file name.
			//
			// Example:
			//    mv "IMG_2818.JPG" "2012-12-21 00:23:16.jpg"
			fmt.Printf("mv %q %q\n", rename.origPath, rename)
		}
	}
}

// Rename contains the original file path and the parsed date.
type Rename struct {
	origPath string
	date     time.Time
	// suffix is used to prevent duplicate file names.
	suffix string
}

// GetRename returns the rename information, based on the embeded date
// information of the provided file.
func GetRename(origPath string) (rename *Rename, err error) {
	// Open file.
	fr, err := os.Open(origPath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	// Read and parse blocks.
	block := make([]byte, 4096)
	for {
		n, err := fr.Read(block)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		date, found := GetDate(block[:n])
		if found {
			rename = &Rename{
				origPath: origPath,
				date:     date,
			}
			return rename, nil
		}
		if n < len(block) {
			break
		}
		// Make sure we don't miss a date between two blocks.
		_, err = fr.Seek(-maxDateLen, os.SEEK_CUR)
		if err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("no valid date representation for %q.", origPath)
}

// String returns the new path based on the old path and the date information.
// The new path will use the same dir and extension as the old path, but the
// file name will be replaced by a date representation. A suffix might be added
// to avoid duplicate file names.
func (rename *Rename) String() string {
	date := rename.date.Format("2006-01-02 - 15.04.05")
	var suffix string
	if len(rename.suffix) > 0 {
		suffix = ", " + rename.suffix
	}
	dir, file := path.Split(rename.origPath)
	ext := strings.ToLower(path.Ext(file))
	newName := date + suffix + ext
	return dir + newName
}

// SanityCheck returns an error if the date is before 1980 or if it is a future
// date.
func (rename *Rename) SanityCheck() (err error) {
	year := rename.date.Year()
	if year < 1980 {
		return fmt.Errorf("ancient date (before 1980): %v.", rename.date)
	}
	if year > time.Now().Year()+1 {
		return fmt.Errorf("future date: %v.", rename.date)
	}
	return nil
}

// SetSuffix sets the suffix to a number from 0001 to 9999. Note, this function
// should only be called to set the suffix of duplicate file names.
func (rename *Rename) SetSuffix(renames []*Rename) {
	var i int
	for _, r := range renames {
		if r.date == rename.date {
			i++
		}
		if r == rename {
			break
		}
	}
	rename.suffix = fmt.Sprintf("%04d", i)
}

// IsDuplicate returns true if the date isn't unique.
func (rename *Rename) IsDuplicate(renames []*Rename) bool {
	var i int
	for _, r := range renames {
		if r.date == rename.date {
			i++
		}
		if i == 2 {
			return true
		}
	}
	return false
}

// SetSuffixes adds a suffix to all non-unique file names.
func SetSuffixes(renames []*Rename) {
	for _, rename := range renames {
		if rename.IsDuplicate(renames) {
			rename.SetSuffix(renames)
		}
	}
}

const maxDateLen = int64(len("2006:01:02 15:04:05"))

// Canon and Nikon date representation.
var regDateColon = regexp.MustCompile("[0-9]{4}:[0-9]{2}:[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}")

// GetDate parses buf looking for various date representations.
func GetDate(buf []byte) (date time.Time, found bool) {
	rawDate := regDateColon.Find(buf)
	if rawDate != nil {
		date, err := time.Parse("2006:01:02 15:04:05", string(rawDate))
		if err == nil {
			return date, true
		}
	}
	return time.Time{}, false
}
