package mongodb

import (
	"testing"

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
		{"21.37", 21.37},
	}

	for _, d := range data {
		f.Add(d.input, d.expected)
	}

	f.Fuzz(func(t *testing.T, value string, expected float64) {
		res := parsePrice(value)
		assert.Equal(t, expected, res)
	})
}
