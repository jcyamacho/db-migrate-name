package generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Name(t *testing.T) {
	t.Parallel()

	now := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	format := "20060102150405"

	tests := []struct {
		name     string
		elems    []string
		gen      *Generator
		expected string
	}{
		{
			name:  "simple words",
			elems: []string{"Hello", "World"},
			gen: New(func(g *Generator) {
				g.TimePrefix = false
			}),
			expected: "Hello_World",
		},
		{
			name:  "simple words to lower",
			elems: []string{"Hello", "World"},
			gen: New(
				func(g *Generator) {
					g.TimePrefix = false
				},
				WithLowerCase(true),
			),
			expected: "hello_world",
		},
		{
			name:  "simple words with date prefix",
			elems: []string{"Hello", "World"},
			gen: New(
				WithTimePrefix(now, format),
				WithSeparator("-"),
			),
			expected: now.Format(format) + "-Hello-World",
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			name := test.gen.Name(test.elems)
			assert.Equal(t, test.expected, name)
		})
	}
}
