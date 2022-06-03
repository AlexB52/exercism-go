package pov

type Tree struct {
	value    string
	children []*Tree
	parent   *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	tree := &Tree{value: value, children: children}
	for _, trc := range tree.Children() {
		trc.parent = tree
	}
	return tree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice ßcontaining the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
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

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	var refTree = map[string]*Tree{}
	BuildRef(refTree, tr)

	return BuildTree(refTree, &Tree{value: from})
}

func BuildRef(m map[string]*Tree, t *Tree) {
	m[t.Value()] = t
	for _, n := range t.Children() {
		m[n.Value()] = n
		BuildRef(m, n)
	}
}

func BuildTree(ref map[string]*Tree, node *Tree) *Tree {
	tree, ok := ref[node.Value()]
	if !ok {
		return nil
	}

	children := tree.Children()
	if tree.parent != nil {
		children = append(children, tree.parent)
	}

	for _, n := range children {
		if node.parent != nil && n.Value() == node.parent.Value() {
			continue
		}
		child := &Tree{n.Value(), nil, node}
		node.children = append(node.children, child)
		BuildTree(ref, child)
	}

	return node
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	return nil
}
