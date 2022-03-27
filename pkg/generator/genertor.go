package generator

import (
	"strings"
	"time"
)

type Generator struct {
	TimePrefix       bool
	TimePrefixFormat string
	Time             time.Time
	Separator        string
	LowerCase        bool
}

func (g *Generator) Name(elems []string) string {
	if g.TimePrefix {
		timePrefix := g.Time.Format(g.TimePrefixFormat)
		elems = append([]string{timePrefix}, elems...)
	}
	name := strings.Join(elems, g.Separator)
	if g.LowerCase {
		name = strings.ToLower(name)
	}
	return name
}

type Config func(*Generator)

func New(configs ...Config) *Generator {
	opt := &Generator{
		TimePrefix:       true,
		TimePrefixFormat: "20060102150405",
		Time:             time.Now(),
		Separator:        "_",
		LowerCase:        false,
	}
	for _, config := range configs {
		config(opt)
	}
	return opt
}

func WithTimePrefix(t time.Time, format string) Config {
	return func(g *Generator) {
		g.TimePrefix = true
		g.Time = t
		g.TimePrefixFormat = format
	}
}

func WithSeparator(separator string) Config {
	return func(g *Generator) {
		g.Separator = separator
	}
}

func WithLowerCase(lowerCase bool) Config {
	return func(g *Generator) {
		g.LowerCase = lowerCase
	}
}
