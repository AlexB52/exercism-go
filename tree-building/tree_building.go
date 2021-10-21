package tree

import (
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, RecordSortingMethod(records))
	return BuildNode(records, &Node{ID: 0}), nil
}

func BuildNode(records []Record, node *Node) *Node {
	for _, record := range records {
		if node.ID == record.ID {
			continue
		}
		if node.ID != record.Parent {
			continue
		}

		newNode := &Node{ID: record.ID}
		BuildNode(records, newNode)
		node.Children = append(node.Children, newNode)
	}

	return node
}

func RecordSortingMethod(records []Record) func(i, j int) bool {
	return func(i, j int) bool {
		return records[i].ID < records[j].ID
	}
}
