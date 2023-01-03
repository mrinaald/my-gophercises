package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mrinaald/my-gophercises/pkg/htmllinkparser"
)

func main() {
	htmlFile := flag.String("html", "", "The input html file to parse")
	flag.Parse()

	fileInfo, err := os.Stat(*htmlFile)
	if err != nil {
		log.Fatalf("Error while opening file: [%v], Error: %v", *htmlFile, err)
	}

	if fileInfo.IsDir() {
		log.Fatalf("Input path [%v] is a directory. Need a file path\n", *htmlFile)
	}

	htmlReader, err := os.Open(*htmlFile)
	if err != nil {
		log.Fatalf("Error while opening file: [%v], Error: %v", *htmlFile, err)
	}
	defer htmlReader.Close()

	links, _ := htmllinkparser.ParseLinks(htmlReader)
	fmt.Printf("%+v\n", links)
}
