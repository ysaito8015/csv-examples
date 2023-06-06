package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	s := `名前,年齢,身長,体重
Tanaka,31,190cm,97kg
Suzuki,46,180cm,79kg
Matsui,45,188cm,95kg
`

	r := csv.NewReader(strings.NewReader(s))

	records, err := r.ReadAll()
	if err != nil {
		if e, ok := err.(*csv.ParseError); ok {
			n := 0
			switch e.Err {
			case csv.ErrBareQuote:
				n = 1
			case csv.ErrQuote:
				n = 2
			case csv.ErrFieldCount:
				n = 3
			}
			log.Fatal("\nError: ", n, "\n", e.Err, "\nStartLine:", e.StartLine, "\nLine:", e.Line, "\nColumn:", e.Column)
		}
		log.Fatal(err)
	}
	fmt.Println(records)
}
