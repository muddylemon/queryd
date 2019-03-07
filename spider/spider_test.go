package spider_test

import (
	"testing"

	"github.com/muddylemon/queryd/spider"
	"github.com/stretchr/testify/assert"
)

func Test_CreateQuestion(t *testing.T) {
	testQuestion := spider.CreateQuestion(55, "The Monkey Sings Tonight")
	//assert.Equal(t, testQuestion.Text, "The Chicken Screams Today")
	assert.Equal(t, testQuestion.Text, "The Monkey Sings Tonight")

}

func Test_ProcessText(t *testing.T) {
	results := spider.Process()
	assert.Nil(t, results)
}
