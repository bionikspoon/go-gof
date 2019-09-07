package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type PolishNotationParser struct {
	symbols []string
}

type ExpressionBase interface {
	Evaluate() int
	String() string
}

func (p *PolishNotationParser) Parse(command string) ExpressionBase {
	p.symbols = strings.Split(command, " ")
	defer func() { p.symbols = []string{} }()
	return p.parseNextExpression()
}

func (p *PolishNotationParser) parseNextExpression() ExpressionBase {
	x := p.symbols[0]
	value, err := strconv.Atoi(x)

	if err != nil {
		return p.parseNonTerminalExpression()
	} else {
		p.symbols = p.symbols[1:]
		return IntegerExpression{value}
	}
}
func (p *PolishNotationParser) parseNonTerminalExpression() ExpressionBase {
	var symbol string
	symbol, p.symbols = p.symbols[0], p.symbols[1:]

	expr1 := p.parseNextExpression()
	expr2 := p.parseNextExpression()

	switch symbol {
	case "+":
		return AdditionExpression{expr1, expr2}
	case "-":
		return SubtractionExpression{expr1, expr2}
	default:
		panic(fmt.Sprintf("Unknown symbol %q", symbol))
	}
}

type IntegerExpression struct {
	value int
}

func (e IntegerExpression) String() string {
	return fmt.Sprint(e.value)
}

func (e IntegerExpression) Evaluate() int {
	return e.value
}

type AdditionExpression struct {
	expr1 ExpressionBase
	expr2 ExpressionBase
}

func (e AdditionExpression) String() string {
	return fmt.Sprintf("(%s + %s)", e.expr1, e.expr2)
}

func (e AdditionExpression) Evaluate() int {
	return e.expr1.Evaluate() + e.expr2.Evaluate()
}

type SubtractionExpression struct {
	expr1 ExpressionBase
	expr2 ExpressionBase
}

func (e SubtractionExpression) String() string {
	return fmt.Sprintf("(%s - %s)", e.expr1, e.expr2)
}

func (e SubtractionExpression) Evaluate() int {
	return e.expr1.Evaluate() - e.expr2.Evaluate()
}
