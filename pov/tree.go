package pov

func New(value string, children ...*Tree) *Tree {
	tree := &Tree{value: value, children: children}
	for _, trc := range tree.Children() {
		trc.parent = tree
	}
	return tree
}

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}
