package expression_test

import (
	"testing"

	"github.com/interactiv/dbal/query/expression"
	"github.com/interactiv/expect"
)

func TestAndX(t *testing.T) {
	e := expect.New(t)
	for _, fixture := range provideDataForAndX() {
		composite := expr().AndX()
		for _, part := range fixture.parts {
			composite.Add(part)
		}
		e.Expect(composite.String()).ToEqual(fixture.expected)
	}
}

type CompositeFixture2 struct {
	parts    []string
	expected string
}

func provideDataForAndX() []CompositeFixture2 {
	return []CompositeFixture2{
		{
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			[]string{"u.user = 1", "u.group_id = 1"},
			"(u.user = 1) AND (u.group_id = 1)",
		},
		{
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			[]string{"u.group_id = 1", "u.group_id = 2"},
			"(u.group_id = 1) AND (u.group_id = 2)",
		},
		{
			[]string{
				"u.user = 1",
				expression.Composite{
					expression.OR,
					[]string{"u.group_id = 1", "u.group_id = 2"},
				}.String(),
			},
			"(u.user = 1) AND ((u.group_id = 1) OR (u.group_id = 2))",
		},
		{
			[]string{
				"u.group_id = 1",
				expression.Composite{
					expression.AND,
					[]string{"u.user = 1", "u.group_id = 2"},
				}.String(),
			},
			"(u.group_id = 1) AND ((u.user = 1) AND (u.group_id = 2))",
		},
	}
}

func TestOrX(t *testing.T) {
	e := expect.New(t)
	for _, fixture := range provideDataForOrX() {
		composite := expr().OrX()
		for _, part := range fixture.parts {
			composite.Add(part)
		}
		e.Expect(fixture.expected).ToEqual(composite.String())

	}

}

func provideDataForOrX() []CompositeFixture2 {
	return []CompositeFixture2{
		{
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			[]string{"u.user = 1", "u.group_id = 1"},
			"(u.user = 1) OR (u.group_id = 1)",
		},
		{
			[]string{"u.user = 1"},
			"u.user = 1",
		},
		{
			[]string{"u.group_id = 1", "u.group_id = 2"},
			"(u.group_id = 1) OR (u.group_id = 2)",
		},
		{
			[]string{
				"u.user = 1",
				expression.Composite{
					expression.OR,
					[]string{"u.group_id = 1", "u.group_id = 2"},
				}.String(),
			},
			"(u.user = 1) OR ((u.group_id = 1) OR (u.group_id = 2))",
		},
		{
			[]string{
				"u.group_id = 1",
				expression.Composite{
					expression.AND,
					[]string{"u.user = 1", "u.group_id = 2"},
				}.String(),
			},
			"(u.group_id = 1) OR ((u.user = 1) AND (u.group_id = 2))",
		},
	}
}

func TestComparison(t *testing.T) {
	e := expect.New(t)
	for _, fixture := range provideDataForComparison() {
		part := expr().Comparison(fixture.x, fixture.Type, fixture.y)
		e.Expect(part).ToEqual(fixture.expected)
	}

}

type ComparisonFixture struct {
	x        string
	Type     expression.TYPE
	y        string
	expected string
}

func provideDataForComparison() []ComparisonFixture {
	return []ComparisonFixture{
		{"u.user_id", expression.EQ, "1", "u.user_id = 1"},
		{"u.user_id", expression.NEQ, "1", "u.user_id <> 1"},
		{"u.salary", expression.LT, "10000", "u.salary < 10000"},
		{"u.salary", expression.LTE, "10000", "u.salary <= 10000"},
		{"u.salary", expression.GT, "10000", "u.salary > 10000"},
		{"u.salary", expression.GTE, "10000", "u.salary >= 10000"},
	}
}

func TestEq(t *testing.T) {
	e := expect.New(t)
	e.Expect("u.user_id = 1").ToEqual(expr().Eq("u.user_id", "1"))
}

func TestNeq(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().Neq("u.user_id", "1")).ToEqual("u.user_id <> 1")
}

func TestLt(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().Lt("u.salary", "10000")).ToEqual("u.salary < 10000")
}

func TestLte(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().Lte("u.salary", "10000")).ToEqual("u.salary <= 10000")
}

func TestGt(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().Gt("u.salary", "10000")).ToEqual("u.salary > 10000")
}

func TestGte(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().Gte("u.salary", "10000")).ToEqual("u.salary >= 10000")
}

func TestIsNull(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().IsNull("u.deleted")).ToEqual("u.deleted IS NULL")
}

func TestIsNotNull(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().IsNotNull("u.updated")).ToEqual("u.updated IS NOT NULL")
}

func TestIn(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().In("u.groups", []interface{}{1, 3, 4, 7}...)).ToEqual("u.groups IN (1, 3, 4, 7)")
}

func TestInWithPlaceholder(t *testing.T) {
	e := expect.New(t)
	e.Expect("u.groups IN (?)").ToBe(expr().In("u.groups", "?"))
}

func TestNotIn(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().NotIn("u.groups", []interface{}{1, 3, 4, 7}...)).ToEqual("u.groups NOT IN (1, 3, 4, 7)")
}

func TestNotInWithPlaceholder(t *testing.T) {
	e := expect.New(t)
	e.Expect(expr().NotIn("u.groups", ":values")).ToEqual("u.groups NOT IN (:values)")
}

func expr() expression.Builder {
	return expression.Builder{(expression.Connection)(nil)}
}
