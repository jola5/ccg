package ngrams

import (
	"bufio"
	"ccg/isdir"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func countRune(c rune) {
	// TODO
}

// CountInFile : Return the sum of individual n-grams in the given file
func CountInFile(filePath string, ngram int) map[string]int {
	isDir, _ := isdir.IsDirectory(filePath)
	if isDir {
		return nil
	}

	fmt.Println(filePath)

	file, e := os.Open(filePath)
	check(e)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			countRune(c)
		}
	}

	check(scanner.Err())

	return nil
}
