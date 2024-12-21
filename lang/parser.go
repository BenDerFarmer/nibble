package fishy

import (
	"fmt"
	"strconv"
)

type ExprType int

const (
	ExprVoid ExprType = iota
	ExprInt
	ExprStr
	ExprVar
	ExprFuncall
)

type Expr struct {
	Type      ExprType
	AsInt     int
	AsStr     string
	AsVar     string
	AsFuncall Funcall
}

type Funcall struct {
	Name string
	Args []Expr
}

type Parser struct {
	tokens []Token
}

func NewParser(tokens *[]Token) Parser {
	return Parser{
		tokens: *tokens,
	}
}

func (p *Parser) ParseAll() []Expr {

	var exprs []Expr

	for cursor := 0; cursor < len(p.tokens); {

		expr, next, err := p.parse(cursor)

		if err != nil {
			panic(err)
		}

		exprs = append(exprs, expr)

		cursor = next

	}

	return exprs
}

func (p *Parser) parse(index int) (Expr, int, error) {
	token := p.tokens[index]

	switch token.TokenType {
	case TokenSym:
		if p.tokens[index+1].TokenType != TokenOParen {
			return Expr{Type: ExprVar, AsVar: string(token.Text)}, index + 1, nil
		} else {

			start := index + 2
			stop := p.findClosingParen(start)

			if stop == -1 {
				return Expr{}, 0, fmt.Errorf("Closing Paren not found for: %s", token)
			}

			args, err := p.parseArgs(start, stop)

			if err != nil {
				return Expr{}, 0, err
			}

			return Expr{Type: ExprFuncall, AsFuncall: Funcall{Name: string(token.Text), Args: args}}, stop + 1, nil

		}
	case TokenNum:

		num, err := strconv.Atoi(string(token.Text))
		if err != nil {
			return Expr{}, 0, err
		}

		return Expr{Type: ExprInt, AsInt: num}, index + 1, nil

	case TokenStr:
		return Expr{Type: ExprStr, AsStr: string(token.Text)}, index + 1, nil

	default:
		return Expr{}, 0, fmt.Errorf("can not parse Tokens(%s) to expr", token)

	}
}
func (p *Parser) parseArgs(start int, stop int) ([]Expr, error) {

	var exprs []Expr

	expectComma := false

	for i := start; i < stop; {
		token := p.tokens[i]

		if expectComma && token.TokenType == TokenComma {
			i++
			expectComma = false
			continue
		}

		if expectComma && token.TokenType != TokenComma {
			return exprs, fmt.Errorf("Expected Comma not found")
		}

		if !expectComma && token.TokenType == TokenComma {
			return exprs, fmt.Errorf("Unexpected Comma in args")
		}

		expr, next, err := p.parse(i)
		if err != nil {
			return exprs, err
		}

		i = next

		exprs = append(exprs, expr)

		expectComma = true

	}

	return exprs, nil
}

func (p *Parser) findClosingParen(start int) int {

	numOfParenToSkip := 0

	for i := start; i < len(p.tokens); i++ {

		token := p.tokens[i]

		if token.TokenType == TokenOParen {
			numOfParenToSkip++
		}

		if token.TokenType != TokenCParen {
			continue
		}

		if numOfParenToSkip == 0 {
			return i
		} else {
			numOfParenToSkip--
		}
	}

	return -1
}
