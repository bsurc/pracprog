package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	// START_OMIT
	var infile *os.File
	var err error
	infile, err = os.Open("myfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	var reader csv.Reader
	reader = csv.NewReader(infile)
	var value string
	var row []string
	for {
		row, err = reader.Read()
		if err != nil {
			break
		}
		// do something with items in rows
		value = row[0] + row[1]
	}
	if err != io.EOF {
		log.Fatal(err)
	}
	infile.Close()
	// END_OMIT
}
