package spider

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
