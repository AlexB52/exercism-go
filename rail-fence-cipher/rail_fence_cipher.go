package railfence

import (
	"sort"
)

type Node struct {
	x, y  int
	value string
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

func Encode(message string, rails int) (result string) {
	var nodes []Node
	projection := DiagonaleProjection(rails)
	for i, r := range message {
		n := projection(i)
		n.value = string(r)
		nodes = append(nodes, n)
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].y < nodes[j].y {
			return true
		} else if nodes[i].y > nodes[j].y {
			return false
		}
		return nodes[i].x < nodes[j].x
	})

	for _, n := range nodes {
		result += n.value
	}
	return result
}

func Decode(message string, rails int) (result string) {
	var nodes = make([]Node, len(message))
	projection := DiagonaleProjection(rails)
	for i := 0; i < len(message); i++ {
		nodes[i] = projection(i)
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].y < nodes[j].y {
			return true
		} else if nodes[i].y > nodes[j].y {
			return false
		}
		return nodes[i].x < nodes[j].x
	})

	for i, n := range nodes {
		nodes[i] = Node{n.x, n.y, string(message[i])}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].x < nodes[j].x
	})

	for _, n := range nodes {
		result += n.value
	}
	return result
}
