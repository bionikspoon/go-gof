package parser

import (
	"fmt"
	"testing"
)

func TestPolishNotationParser(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{"+ 5 6", "(5 + 6) = 11"},
		{"- 6 5", "(6 - 5) = 1"},
		{"+ - 4 5 6", "((4 - 5) + 6) = 5"},
		{"+ 4 - 5 6", "(4 + (5 - 6)) = 3"},
		{"+ - + - - 2 3 4 + - -5 6 + -7 8 9 10", "(((((2 - 3) - 4) + ((-5 - 6) + (-7 + 8))) - 9) + 10) = -14"},
	}

	parser := PolishNotationParser{}
	for _, tt := range cases {
		desc := fmt.Sprintf("given %q", tt.command)
		t.Run(desc, func(t *testing.T) {
			expression := parser.Parse(tt.command)

			if got := fmt.Sprintf("%s = %d", expression, expression.Evaluate()); got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
