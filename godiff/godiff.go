package main

import "bytes"
import "flag"
import "io/ioutil"
import "log"
import "os"
import "os/exec"
import "path"
import "strings"

func main() {
	flag.Parse()
	for _, srcPath := range flag.Args() {
		err := godiff(srcPath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func godiff(srcPath string) (err error) {
	if path.Ext(srcPath) != ".go" {
		return nil
	}
	cmd := exec.Command("gofmt", "-d=true", srcPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		if err.Error() != "exit status 2" {
			return err
		}
	}
	if out.Len() > 0 {
		diffPath, err := mktemp(out.Bytes())
		if err != nil {
			return err
		}
		err = view(diffPath)
		if err != nil {
			return err
		}
		err = os.Remove(diffPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func mktemp(buf []byte) (tmpPath string, err error) {
	cmd := exec.Command("mktemp", "-u")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	tmpPath = strings.TrimSpace(out.String()) + ".patch"
	err = ioutil.WriteFile(tmpPath, buf, 0644)
	if err != nil {
		return "", err
	}
	return tmpPath, nil
}

func view(diffPath string) (err error) {
	cmd := exec.Command("geany", "-i", diffPath)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
