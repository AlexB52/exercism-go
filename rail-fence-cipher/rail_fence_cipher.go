package railfence

import (
	"sort"
)

type Node struct {
	x, y  int
	value string
}

func Encode(message string, rails int) (result string) {
	var nodes = BuildNodes(len(message), rails)
	return Code(message, nodes, DiagonalSorting(nodes), RailsSorting(nodes))
}

func Decode(message string, rails int) (result string) {
	var nodes = BuildNodes(len(message), rails)
	return Code(message, nodes, RailsSorting(nodes), DiagonalSorting(nodes))
}

func Code(message string, nodes []Node, sort1, sort2 func(int, int) bool) (result string) {
	sort.Slice(nodes, sort1)

	for i, n := range nodes {
		nodes[i] = Node{n.x, n.y, string(message[i])}
	}

	sort.Slice(nodes, sort2)

	for _, n := range nodes {
		result += n.value
	}
	return result
}

func BuildNodes(length, rails int) []Node {
	var nodes = make([]Node, length)
	projection := DiagonaleProjection(rails)
	for i := 0; i < length; i++ {
		nodes[i] = projection(i)
	}
	return nodes
}

func DiagonaleProjection(rail int) func(int) Node {
	return func(pos int) Node {
		x, y := pos, pos%(2*(rail-1))
		if y >= rail {
			y = rail - (y % rail) - 2
		}
		return Node{x: x, y: y}
	}
}

func RailsSorting(nodes []Node) func(int, int) bool {
	return func(i, j int) bool {
		if nodes[i].y < nodes[j].y {
			return true
		} else if nodes[i].y > nodes[j].y {
			return false
		}
		return nodes[i].x < nodes[j].x
	}
}

func DiagonalSorting(nodes []Node) func(int, int) bool {
	return func(i, j int) bool {
		return nodes[i].x < nodes[j].x
	}
}
