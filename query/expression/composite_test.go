package expression_test

import (
	"testing"

	"github.com/interactiv/dbal/query/expression"
	"github.com/interactiv/expect"
)

func TestCompositeExpression(t *testing.T) {
	e := expect.New(t)
	expr := &expression.Composite{expression.OR, []string{"u.group_id = 1"}}
	e.Expect(expr.Length()).ToBe(1)
	expr.Add("u.group_id = 2")
	e.Expect(expr.Length()).ToBe(2)
}

func TestCompositeUsageAndGeneration(t *testing.T) {

	for _, fixture := range testData() {
		expr := expression.Composite{fixture.TYPE, fixture.parts}
		expect.Expect(expr.String(), t).ToBe(fixture.expected)
	}
}

type CompositeFixture struct {
	expression.TYPE
	parts    []string
	expected string
}

func testData() []CompositeFixture {

	return []CompositeFixture{
		{
			expression.AND,
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			expression.AND,
			[]string{"u.user = 1", "u.group_id = 1"},
			"(u.user = 1) AND (u.group_id = 1)",
		},
		{
			expression.OR,
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			expression.OR,
			[]string{"u.group_id = 1", "u.group_id = 2"},
			"(u.group_id = 1) OR (u.group_id = 2)",
		},
		{
			expression.AND,
			[]string{
				"u.user = 1",
				(expression.Composite{
					expression.OR,
					[]string{"u.group_id = 1", "u.group_id = 2"},
				}).String(),
			},
			"(u.user = 1) AND ((u.group_id = 1) OR (u.group_id = 2))",
		},
		{
			expression.OR,
			[]string{
				"u.group_id = 1",

				(expression.Composite{
					expression.AND,
					[]string{"u.user = 1", "u.group_id = 2"},
				}).String(),
			},
			"(u.group_id = 1) OR ((u.user = 1) AND (u.group_id = 2))",
		},
	}
}
