package spider_test

import (
	"testing"

	"github.com/muddylemon/queryd/spider"
	"github.com/stretchr/testify/assert"
)

func Test_CreateQuestion(t *testing.T) {
	testQuestion := spider.CreateQuestion("The Monkey Sings Tonight", "Cool")
	//assert.Equal(t, testQuestion.Text, "The Chicken Screams Today")
	assert.Equal(t, testQuestion.Title, "The Monkey Sings Tonight")

}

func Test_ProcessText(t *testing.T) {
	results, err := spider.Process()
	assert.Nil(t, err)
	assert.Equal(t, []*spider.Question{
		&spider.Question{
			Title:   "Hello",
			Details: "Hogwarts",
		},
	}, results)
}
