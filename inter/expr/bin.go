package expr

import "dl/token"

type Bin struct {
	Expr
	expr1 *Expr
	expr2 *Expr
}

func NewBin(op *token.Token, e1, e2 *Expr) *Bin {
	bin := &Bin{
		expr1: e1,
		expr2: e2,
	}
	bin.op = op
	bin.addChild(e1)
	bin.addChild(e2)
	return bin
}

func (b *Bin) addChild(n *Expr) {
	b.Expr.addChild(n)
}
