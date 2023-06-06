package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/gocarina/gocsv"
)

type Client struct {
	Id   string `csv:"id"`
	Name string `csv:"名前"`
	Age  string `csv:"年齢"`
}

func main() {
	file, err := os.Open("export-sjis.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	clients := []*Client{}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
		return r
	})

	if err := gocsv.UnmarshalFile(file, &clients); err != nil {
		log.Fatal(err)
	}

	for _, client := range clients {
		fmt.Println(client.Name)
	}
}
