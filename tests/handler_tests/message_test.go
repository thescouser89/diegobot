package handler_tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/thescouser89/diegobot/handlers"
	"testing"
)

func TestExtractLink(t *testing.T) {
	assert := assert.New(t)
	link1 := "ksj ksjdf sldk http://haha.com"
	assert.Equal(handlers.ExtractLink(link1), "http://haha.com")
}
