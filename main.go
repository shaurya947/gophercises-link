package main

import (
	"flag"
	"fmt"
	"gophercises-link/linkparser"
	"log"
	"os"
)

func main() {
	filenamePtr := flag.String("file", "ex1.html", "HTML file to be parsed")
	flag.Parse()

	filePtr, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatalf("Error opening file %s: %+v", *filenamePtr, err)
	}

	links, err := linkparser.ParseFromReader(filePtr)
	if err != nil {
		log.Fatalf("Error parsing links from file %s: %+v", *filenamePtr, err)
	}

	fmt.Printf("%+v\n", links)
}
