package tree

import (
	"errors"
	"fmt"
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
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	if len(records) == 0 {
		return nil, nil
	}

	if ok, err := ValidateRootRecord(records[0]); !ok {
		return nil, err
	}

	nodes := make([]*Node, len(records))
	nodes[0] = &Node{ID: 0}

	for i := 1; i < len(records); i++ {
		if ok, err := ValidateRecord(i, records[i]); !ok {
			return nil, err
		}

		nodes[i] = &Node{ID: records[i].ID}
		parent := nodes[records[i].Parent]
		parent.Children = append(parent.Children, nodes[i])
	}

	return nodes[0], nil
}

func ValidateRootRecord(record Record) (bool, error) {
	if record.Parent != 0 {
		return false, errors.New("root node has parent")
	}

	if record.ID != 0 {
		return false, errors.New("no root node")
	}

	return true, nil
}

func ValidateRecord(index int, record Record) (bool, error) {
	if index != record.ID {
		return false, fmt.Errorf("non-contiguous record ID: %+v", record)
	}

	if record.Parent != 0 && record.ID <= record.Parent {
		return false, fmt.Errorf("higher parent id than record id: %+v", record)
	}

	return true, nil
}
