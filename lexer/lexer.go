package lexer

import (
	"bufio"
	"dl/tag"
	"dl/token"
	"io"
	"os"
	"unicode"
)

var err error

type Lexer struct {
	reader   bufio.Reader
	line     int
	peek     byte
	keywords map[string]tag.Tag
}

func New(file *os.File) *Lexer {
	return &Lexer{
		reader: *bufio.NewReader(file),
		peek:   ' ',
		line:   1,
		keywords: map[string]tag.Tag{
			"programa": tag.PROGRAM,
			"inicio":   tag.BEGIN,
			"fim":      tag.END,
			"se":       tag.IF,
			"escreva":  tag.WRITE,
			"inteiro":  tag.INT,
		},
	}
}

func (l *Lexer) Line() int {
	return l.line
}

func (l *Lexer) nextChar() {

	if l.peek == '\n' {
		l.line = l.line + 1
	}

	l.peek, err = l.reader.ReadByte()
}

func (l *Lexer) isWhiteSpace(b byte) bool {
	switch string(b) {
	case " ":
		return true
	case "\t":
		return true
	case "\n":
		return true
	case "\r":
		return true
	default:
		return false
	}
}

func isIdStart(c byte) bool {
	return (unicode.IsLetter(rune(c)) || c == '_')
}

func isIdPart(c byte) bool {
	return (isIdStart(c) || unicode.IsDigit(rune(c)))
}

func (l *Lexer) NextToken() *token.Token {

	if err != nil {
		if err == io.EOF {
			return token.New(tag.EOF, "")
		}
		panic(err)
	}

	for l.isWhiteSpace(l.peek) {
		l.nextChar()
	}

	switch l.peek {
	case '=':
		l.nextChar()
		return token.New(tag.ASSIGN, "=")
	case '+':
		l.nextChar()
		return token.New(tag.SUM, "+")
	case '-':
		l.nextChar()
		return token.New(tag.SUB, "-")
	case '*':
		l.nextChar()
		return token.New(tag.MUL, "*")
	case ';':
		l.nextChar()
		return token.New(tag.SEMI, ";")
	case '|':
		l.nextChar()
		return token.New(tag.OR, "|")
	case '(':
		l.nextChar()
		return token.New(tag.LPAREN, "(")
	case ')':
		l.nextChar()
		return token.New(tag.RPAREN, ")")
	case '<':
		l.nextChar()
		if l.peek == '=' {
			l.nextChar()
			return token.New(tag.LE, "<=")
		}
		return token.New(tag.LT, "<")
	case '>':
		l.nextChar()
		return token.New(tag.GT, ">")
	case '.':
		l.nextChar()
		return token.New(tag.DOT, ".")
	default:
		if unicode.IsDigit(rune(l.peek)) {
			number := ""
			for unicode.IsDigit(rune(l.peek)) {
				number = number + string(l.peek)
				l.nextChar()

				if l.peek == '.' {
					number = number + string(l.peek)
					l.nextChar()
					for unicode.IsDigit(rune(l.peek)) {
						number = number + string(l.peek)
						l.nextChar()
					}
					return token.New(tag.LIT_REAL, number)
				}
			}
			return token.New(tag.LIT_INT, number)
		} else if isIdStart(l.peek) {
			id := ""
			for {
				id = id + string(l.peek)
				l.nextChar()
				if !isIdPart(l.peek) {
					break
				}
			}
			if _, ok := l.keywords[id]; ok {
				return token.New(l.keywords[id], id)
			}
			return token.New(tag.ID, id)
		}
	}
	unk := string(l.peek)
	l.nextChar()
	return token.New(tag.UNK, unk)
}
