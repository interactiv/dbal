package expression

type TYPE string

const (
	OR  TYPE = "OR"
	AND TYPE = "AND"
)

type CompositeExpression struct {
	Type  TYPE
	Parts []string
}

func (ce CompositeExpression) Length() int {
	return len(ce.Parts)
}

func (ce *CompositeExpression) Add(expression string) {
	ce.Parts = append(ce.Parts, expression)
}
