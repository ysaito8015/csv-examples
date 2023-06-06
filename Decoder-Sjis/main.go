package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	f, err := os.Open("file-sjis.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
