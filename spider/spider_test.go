package spider_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateQuestion(t *testing.T) {
	testQuestion := queryd.CreateQuestion(55, "The Monkey Sings Tonight")

	assert.Equal(t, testQuestion.Text, "The Chicken Screams Today")

}
