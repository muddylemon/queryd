package spider

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	prose "gopkg.in/jdkato/prose.v2"
)

// Question is a question
type Question struct {
	gorm.Model
	Title   string `gorm:"type:varchar(250);unique_index"`
	Details string `gorm:"type:text(16)"`
	Link    string
}

type QuestionEntity struct {
	gorm.Model
	QuestionID uint
	Question   Question
	Name       string
	Type       string
}

// CreateQuestion creates a question
func CreateQuestion(title, details string) *Question {
	return &Question{
		Title:   title,
		Details: details,
	}
}

// Process processes the process
func Process() ([]*Question, error) {
	var counter int
	var questions []*Question

	db := createDB()

	defer db.Close()

	db.AutoMigrate(&Question{})
	db.AutoMigrate(&QuestionEntity{})

	csvFile, _ := os.Open("../data/qd-raw.csv")
	r := csv.NewReader(csvFile)
	defer func() {
		log.Println("Done")
	}()
	for {
		record := read(r)
		if counter > 1000 {
			log.Println("OMG slow down")
			break
		}
		q := CreateQuestion(record[0], record[1])
		db.Create(&q)

		doc, err := prose.NewDocument(fmt.Sprintf("<h1>%s</h1> <article>%s</article>", record[0], record[1]))
		if err != nil {
			fmt.Printf("Error %+v", err)
		} else {
			for _, ent := range doc.Entities() {
				e := &QuestionEntity{
					QuestionID: q.ID,
					Name:       ent.Text,
					Type:       ent.Label,
				}
				db.Create(&e)
				fmt.Printf("QE: %+v\n\n", e)
				// Lebron James PERSON
				// Los Angeles GPE
			}
		}

		counter++
	}
	return questions, nil
}

func createDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/queryd?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func read(r *csv.Reader) []string {
	record, err := r.Read()

	if err == io.EOF {
		log.Println("Ran out of text yo")
		return nil
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return record
}

// func PageMaker(title, body, slug string) error {

// 	return nil
// }
