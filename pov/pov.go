package pov

import (
	"fmt"
	"regexp"
	"strings"
)

type Tree struct {
	value    string
	children []*Tree
	parent   *Tree
}

func (tr *Tree) PathTo(from, to string) []string {
	re := regexp.MustCompile(fmt.Sprintf(`.*\b%s\b`, to))
	match := re.FindString(tr.FromPov(from).String())

	var result []string
	for _, word := range strings.Split(match, " ") {
		if word == to {
			result = append(result, word)
		}

		if strings.HasPrefix(word, "(") {
			result = append(result, word[1:])
		}

		if strings.HasSuffix(word, ")") {
			result = result[:len(result)-1]
		}
	}

	return result
}

func (tr *Tree) FromPov(from string) *Tree {
	var refTree = map[string]*Tree{}
	MapReferenceTree(refTree, tr)
	return BuildTree(refTree, &Tree{value: from})
}

func MapReferenceTree(m map[string]*Tree, t *Tree) {
	m[t.Value()] = t
	for _, n := range t.Children() {
		m[n.Value()] = n
		MapReferenceTree(m, n)
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
