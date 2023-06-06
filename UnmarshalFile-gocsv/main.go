package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type NotUsed struct {
	Name string
}

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id            string  `csv:"client_id"`
	Name          string  `csv:"client_name"`
	Age           string  `csv:"client_age"`
	NotUsedString string  `csv:"-"`
	NotUsedStruct NotUsed `csv:"-"`
}

func main() {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil {
		log.Fatal(err)
	}

	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		log.Fatal(err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"})
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})
	csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(csvContent) // Display all clients as CSV string
}
