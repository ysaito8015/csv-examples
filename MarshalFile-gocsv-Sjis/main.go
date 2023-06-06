package main

import (
	"encoding/csv"
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
	clients := []*Client{}
	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"})
	clients = append(clients, &Client{Id: "13", Name: "Fred"})

	file, err := os.OpenFile("export-sjis.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// set export encoding as sjis
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()))
		return gocsv.NewSafeCSVWriter(writer)
	})

	gocsv.MarshalFile(&clients, file)
}
