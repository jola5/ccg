package main

import (
	"ccg/ngrams"
	"fmt"
	"github.com/ryanuber/go-glob"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

// @TODO: For dev only
/*
import (
	"github.com/davecgh/go-spew/spew"
)
*/

var (
	app   = kingpin.New("ccg", "character counter greatness - a tool to count characters in files")
	paths = app.Arg("path", "Path to directory or file to include in character counting").Strings()
	globs = app.Flag("globs", "List of file globbing patterns to exclusively include").Default("*").Strings()
	ngram = app.Flag("ngram", "Number of chars to count n-grams instead of single letter chars").Default("1").Int()
)

func countCharsInFile(path string, f os.FileInfo, err error) error {
	for _, pattern := range *globs {
		if glob.Glob(pattern, path) {
			ngrams.CountInFile(path, *ngram)
		}
	}
	return nil
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// defaulting to current working directory
	if len(*paths) == 0 {
		*paths = append(*paths, pwd)
	}

	fmt.Println("#paths : ", len(*paths))
	fmt.Println("paths  : ", *paths)
	fmt.Println("#globs : ", len(*globs))
	fmt.Println("globs  : ", *globs)
	fmt.Println("ngram  : ", *ngram)

	for _, path := range *paths {
		err := filepath.Walk(path, countCharsInFile)
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
