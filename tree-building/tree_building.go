package tree

import (
	"errors"
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
	sort.Slice(records, RecordSortingMethod(records))

	if len(records) == 0 {
		return nil, nil
	}

	err := ValidateTree(records)
	if err != nil {
		return nil, err
	}

	return BuildNode(records, &Node{ID: 0}), nil
}

func ValidateTree(records []Record) error {
	if records[0].Parent != 0 {
		return errors.New("root node has parent")
	}

	if records[0].ID != 0 {
		return errors.New("no root node")
	}

	duplicates := make(map[int]int)
	for i, record := range records {
		if i != record.ID {
			return errors.New("non-continuous")
		}

		if record.Parent != 0 && record.ID <= record.Parent {
			return errors.New("higher id parent of lower id")
		}

		duplicates[record.ID]++
		if duplicates[record.ID] > 1 {
			return errors.New("duplicate node")
		}
	}

	return nil
}

func BuildNode(records []Record, node *Node) *Node {
	for _, record := range records {
		if node.ID == record.ID {
			continue
		}

		if node.ID != record.Parent {
			continue
		}

		node.Children = append(node.Children, BuildNode(records, &Node{ID: record.ID}))
	}

	return node
}

func RecordSortingMethod(records []Record) func(i, j int) bool {
	return func(i, j int) bool {
		return records[i].ID < records[j].ID
	}
}
