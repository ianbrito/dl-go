package token

import (
	"dl/lexer/tag"
	"fmt"
)

type Token struct {
	tag    tag.Tag
	lexeme string
}

func New(tag tag.Tag, lexeme string) *Token {
	return &Token{
		tag:    tag,
		lexeme: lexeme,
	}
}

func (t *Token) Tag() tag.Tag {
	return t.tag
}

func (t *Token) Lexeme() string {
	return t.lexeme
}

func (t *Token) String() string {
	return fmt.Sprintf("<%s, '%s'>", t.tag, t.lexeme)
}
