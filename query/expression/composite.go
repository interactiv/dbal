package expression

import (
	"strings"
)

type TYPE string

const (
	OR  TYPE = "OR"
	AND TYPE = "AND"
)

type Composite struct {
	Type  TYPE
	Parts []string
}

func (ce Composite) Length() int {
	return len(ce.Parts)
}

func (ce *Composite) Add(expression string) {
	ce.Parts = append(ce.Parts, expression)
}

func (ce Composite) String() string {
	switch len(ce.Parts) {
	case 0:
		return ""
	case 1:
		return ce.Parts[0]
	default:
		return "(" + strings.Join(ce.Parts, ") "+string(ce.Type)+" (") + ")"
	}
}
