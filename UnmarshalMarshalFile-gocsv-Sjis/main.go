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
	file, err := os.Open("import-sjis.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	clients := []*Client{}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(transform.NewReader(in, japanese.ShiftJIS.NewDecoder()))
		return r
	})

	if err := gocsv.UnmarshalFile(file, &clients); err != nil {
		log.Fatal(err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "23"})
	clients = append(clients, &Client{Id: "13", Name: "Fred", Age: "0"})

	for _, v := range clients {
		fmt.Println(v)
	}

	exportFile, err := os.OpenFile("export-sjis.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer exportFile.Close()

	// set export encoding as sjis
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()))
		return gocsv.NewSafeCSVWriter(writer)
	})

	gocsv.MarshalFile(&clients, exportFile)
}
