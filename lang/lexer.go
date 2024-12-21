package fishy

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	TokenInvalid TokenType = iota
	TokenSym
	TokenNum
	TokenStr
	TokenOParen
	TokenCParen
	TokenComma
)

type Token struct {
	TokenType TokenType
	Text      []rune
}

type Lexer struct {
	Content []rune
}

func NewLexer(content []rune) Lexer {
	return Lexer{
		Content: content,
	}
}

const (
	StringStartRune = '"'
	StringEndRune   = '\''
)

func (l *Lexer) Lex() (*[]Token, error) {

	var tokens []Token

	for cursor := 0; cursor < len(l.Content); cursor++ {

		rune := l.Content[cursor]

		switch rune {
		case '(':
			tokens = append(tokens, Token{TokenType: TokenOParen})
		case ')':
			tokens = append(tokens, Token{TokenType: TokenCParen})
		case ',':
			tokens = append(tokens, Token{TokenType: TokenComma})
		case StringStartRune:
			next := l.findStringEnd(cursor + 1)
			if next == 0 {
				return &tokens, fmt.Errorf("End of String not found.")
			}

			tokens = append(tokens, Token{TokenType: TokenStr, Text: l.Content[cursor+1 : next]})

			cursor = next
		}

		if unicode.IsLetter(rune) {

			next := l.findSymboleEnd(cursor + 1)

			if next == 0 {
				next = len(l.Content)
			}

			tokens = append(tokens, Token{TokenType: TokenSym, Text: l.Content[cursor:next]})

			cursor = next - 1

		}

		if unicode.IsNumber(rune) {

			next := l.findIntEnd(cursor + 1)

			if next == 0 {
				next = len(l.Content)
			}

			tokens = append(tokens, Token{TokenType: TokenNum, Text: l.Content[cursor:next]})

			cursor = next - 1
		}
	}

	return &tokens, nil
}

func (l *Lexer) findStringEnd(start int) int {

	numOfEndsToSkip := 0

	for i := start; i < len(l.Content); i++ {
		if l.Content[i] == StringEndRune {

			if numOfEndsToSkip == 0 {
				return i
			} else {
				numOfEndsToSkip--
			}
		}

		if l.Content[i] == StringStartRune {
			numOfEndsToSkip++
		}
	}

	return 0
}

func (l *Lexer) findIntEnd(start int) int {

	for i := start; i < len(l.Content); i++ {

		rune := l.Content[i]

		if unicode.IsSpace(rune) || rune == '(' || rune == ')' || rune == ',' || rune == StringStartRune {
			return i
		}
	}

	return len(l.Content)
}

func (l *Lexer) findSymboleEnd(start int) int {

	for i := start; i < len(l.Content); i++ {

		rune := l.Content[i]

		if unicode.IsSpace(rune) || rune == '(' || rune == ')' || rune == ',' || rune == StringStartRune {
			return i
		}

	}

	return len(l.Content)
}
