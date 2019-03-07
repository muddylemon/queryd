package spider

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	prose "gopkg.in/jdkato/prose.v2"
)

// Question is a question
type Question struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// CreateQuestion creates a question
func CreateQuestion(id int, text string) *Question {
	return &Question{
		ID:   id,
		Text: text,
	}
}

// Process processes the process
func Process() error {
	var counter int
	csvFile, _ := os.Open("../data/qd-raw.csv")
	r := csv.NewReader(csvFile)
	defer func() {
		log.Println("Done")
	}()
	for {
		record, err := r.Read()

		if err == io.EOF {
			log.Println("Ran out of text yo")
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		if counter > 10 {
			log.Println("OMG slow down")
			return nil
		}
		doc, _ := prose.NewDocument(fmt.Sprintf("<h1>%s</h1> <article>%s</article>", record[0], record[1]))
		log.Printf("doc: %+v \n\n", doc)
		for _, ent := range doc.Entities() {
			fmt.Println(ent.Text, ent.Label)
			// Lebron James PERSON
			// Los Angeles GPE
		}
		counter++
	}
}
