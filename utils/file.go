package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetFileContents(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)

	return stringContent
}

func GetFileLines(filename string) []string {
	contents := GetFileContents(filename)
	lines := strings.Split(contents, "\n")
	/*
		This won't work: println("lines:", lines)
		When printing slices, Go shows the type and memory address instead of the contents.
		To print the contents of a slice, you need to iterate over it or use a formatting verb that handles slices appropriately.
		fmt.Printf("lines: %s\n", lines)
	*/

	return lines
}
