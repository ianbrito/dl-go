package expr

import (
	"dl/tag"
	"dl/token"
)

type Node interface {
	addChild(n Node)
}

type Expr struct {
	children []Node
	op       *token.Token
	Type     tag.Tag
}

func New(op *token.Token, t tag.Tag) *Expr {
	return &Expr{
		op:   op,
		Type: t,
	}
}

func (e *Expr) addChild(n Node) {
	e.children = append(e.children, n)
}

func (e *Expr) Op() *token.Token {
	return e.op
}

func (e *Expr) GetType() tag.Tag {
	return e.Type
}

func (e *Expr) String() string {
	return e.op.Tag().String()
}
