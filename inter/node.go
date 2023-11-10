package inter

type INode interface {
	addChild()
	getChildren()
}

type Node struct {
	children []*Node
}

func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) getChildren() []*Node {
	return n.children
}
