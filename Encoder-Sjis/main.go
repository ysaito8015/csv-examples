package main

import (
	"encoding/csv"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	records := [][]string{
		[]string{"名前", "年齢", "身長", "体重"},
		[]string{"田中", "31", "190cm", "97kg"},
		[]string{"鈴木", "46", "180cm", "66kg"},
		[]string{"須藤", "45", "188cm", "100kg"},
	}

	f, err := os.Create("sjis-file.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(transform.NewWriter(f, japanese.ShiftJIS.NewEncoder()))

	w.WriteAll(records)
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
