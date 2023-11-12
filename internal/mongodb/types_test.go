package mongodb

import (
	"testing"

	"github.com/dominikus1993/xbox-promotion-checker-app/pkg/games"
	"github.com/stretchr/testify/assert"
)

func FuzzParsePrice(f *testing.F) {
	data := []struct {
		input    string
		expected float64
	}{
		{"", 0},
		{"0", 0},
		{"0.0", 0},
		{"2.1", 2.1},
		{"2.13", 2.13},
		{"21.37", 21.37},
		{"21.37.0", 0},
	}

	for _, d := range data {
		f.Add(d.input, d.expected)
	}

	f.Fuzz(func(t *testing.T, value string, expected float64) {
		res := parsePrice(value)
		assert.Equal(t, expected, res)
	})
}

func TestToXboxGame(t *testing.T) {
	element := screapedXboxHame{
		ID:        "123",
		Title:     "title",
		Link:      "link",
		Price:     "1.23",
		OldPrice:  "2.34",
		CrawledAt: 123,
		Promotion: "10%",
	}

	expected := games.XboxStoreGame{
		ID:       "123",
		Title:    "title",
		Link:     "link",
		Price:    1.23,
		OldPrice: 2.34,
	}

	res := toXboxGame(element)

	assert.Equal(t, expected, res)

}
