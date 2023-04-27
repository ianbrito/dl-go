package parser

import (
	"dl/lexer"
	"dl/tag"
	"dl/token"
	"fmt"
	"os"
)

type Parser struct {
	lexer *lexer.Lexer
	look  *token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
		look:  lexer.NextToken(),
	}
}

func (p *Parser) Move() *token.Token {
	save := p.look
	p.look = p.lexer.NextToken()
	return save
}

func (p *Parser) move() *token.Token {
	save := p.look
	p.look = p.lexer.NextToken()
	return save
}

func (p *Parser) error(s string) {
	fmt.Printf("linha %d : %s", p.lexer.Line(), s)
	os.Exit(0)
}

func (p *Parser) match(t tag.Tag) *token.Token {
	if p.look.Tag() == t {
		return p.move()
	}
	fmt.Print(p.look.Tag())
	fmt.Print(t)
	p.error("Símbolo inesperado")
	return nil
}

func (p *Parser) Parse() {
	p.program()
}

func (p *Parser) program() {
	p.match(tag.PROGRAM)
	p.match(tag.ID)
	p.block()
	p.match(tag.DOT)
	p.match(tag.EOF)
}

func (p *Parser) block() {
	p.match(tag.BEGIN)
	for p.look.Tag() != tag.END {
		p.stmt()
		p.match(tag.SEMI)
	}
	p.match(tag.END)
}

func (p *Parser) stmt() {
	switch p.look.Tag() {
	case tag.BEGIN:
		p.block()
	case tag.INT:
		p.decl()
	case tag.REAL:
		p.decl()
	case tag.BOOL:
		p.decl()
	case tag.ID:
		p.assign()
	case tag.IF:
		p.ifStmt()
	case tag.WRITE:
		p.writeStmt()
	default:
		fmt.Print(p.look.Tag())
		p.error("comando inválido")
	}
}

func (p *Parser) decl() {
	p.move()
	p.match(tag.ID)
}

func (p *Parser) assign() {
	p.match(tag.ID)
	p.match(tag.ASSIGN)
	p.expr()
}

func (p *Parser) expr() {
	p.rel()
	for p.look.Tag() == tag.OR {
		p.move()
		p.rel()
	}
}

func (p *Parser) rel() {
	p.arith()
	for p.look.Tag() == tag.LT || p.look.Tag() == tag.LE || p.look.Tag() == tag.GT {
		p.move()
		p.arith()
	}
}

func (p *Parser) arith() {
	p.term()
	for p.look.Tag() == tag.SUM || p.look.Tag() == tag.SUB {
		p.move()
		p.term()
	}
}

func (p *Parser) term() {
	p.factor()
	for p.look.Tag() == tag.MUL {
		p.move()
		p.factor()
	}
}

func (p *Parser) factor() {
	switch p.look.Tag() {
	case tag.LPAREN:
		p.expr()
		p.match(tag.RPAREN)
	case tag.LIT_INT:
		p.move()
	case tag.LIT_REAL:
		p.move()
	case tag.TRUE:
	case tag.FALSE:
		p.move()
	case tag.ID:
		p.match(tag.ID)
	default:
		p.error("expressão inválida")
	}
}

func (p *Parser) ifStmt() {
	p.match(tag.IF)
	p.match(tag.LPAREN)
	p.expr()
	p.match(tag.RPAREN)
	p.stmt()
}

func (p *Parser) writeStmt() {
	p.move()
	p.match(tag.LPAREN)
	p.match(tag.ID)
	p.match(tag.RPAREN)
}
