package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
}

func workingDirectory() string {
	if *fromDir != "" {
		return *fromDir
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
