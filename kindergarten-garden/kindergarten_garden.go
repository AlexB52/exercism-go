package kindergarten

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var PLANTS = map[byte]string{
	'R': "radishes",
	'V': "violets",
	'C': "clover",
	'G': "grass",
}

type Garden struct {
	PlantsByChidren map[string][]string
}

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	plants, ok = g.PlantsByChidren[child]
	return plants, ok
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	var sortedChildren = make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	validationRe := regexp.MustCompile(fmt.Sprintf("(\n[RVCG]{%d}){2}", 2*len(children)))
	if !validationRe.Match([]byte(diagram)) {
		return nil, errors.New("wrong diagram format")
	}

	// Make diagram as one line once validated
	diagram = strings.ReplaceAll(diagram, "\n", "")

	var plants = map[string][]string{}
	for i, child := range sortedChildren {
		if _, ok := plants[child]; ok {
			return nil, errors.New("duplicate user")
		}

		a := 2 * i
		b := a + 1
		c := 2*len(children) + a
		d := c + 1

		plants[child] = append(plants[child], PLANTS[diagram[a]])
		plants[child] = append(plants[child], PLANTS[diagram[b]])
		plants[child] = append(plants[child], PLANTS[diagram[c]])
		plants[child] = append(plants[child], PLANTS[diagram[d]])
	}

	return &Garden{PlantsByChidren: plants}, nil
}
