package db_test

import (
	"crawler/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieFill(t *testing.T) {
	a := assert.New(t)

	movie := &db.Movie{}
	data := map[string][]string{
		"Title": []string{"标题"},
		"Types": []string{"恐怖", "悬疑"},
	}
	movie.Fill(data)
	a.Equal("标题", movie.Title)
	a.Equal([]string{"恐怖", "悬疑"}, movie.Types)
}
