package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	debug      *bool
	extensions *string
	fromDir    *string
	recur      *bool
	toDir      *string
)

func init() {
	debug = flag.Bool("debug", false, "enable debug mode")
	extensions = flag.String("ext", "", "a comma separated values indicating the file extensions to look for")
	fromDir = flag.String("dir", "", "the directory containing the files")
	recur = flag.Bool("r", false, "recursive finding of files")
	toDir = flag.String("out", "to", "the output directory")
}

func main() {

	flag.Parse()

	dir := workingDirectory()

	fmt.Println(dir)
}

func workingDirectory() string {

	// Get the current directory. This can be useful even when the
	// user provided a working directory, as there is a chance it
	// isn't an absolute path.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// User provided a working directory.
	if *fromDir != "" {

		// Absolute path.
		// TODO Possible issues in Windows.
		if strings.HasPrefix(*fromDir, "/") {
			return *fromDir
		}

		// Relative path.
		return dir + "/" + *fromDir
	}

	// Default to working directory.
	return dir
}
