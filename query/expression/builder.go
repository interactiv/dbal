package expression

import (
	"fmt"
)

// Connection is a database connection
type Connection interface {
	Quote(interface{}, string) string
}

// Expression types
const (
	EQ      TYPE = "="
	NEQ     TYPE = "<>"
	LT      TYPE = "<"
	LTE     TYPE = "<="
	GT      TYPE = ">"
	GTE     TYPE = ">="
	IN      TYPE = "IN"
	NOTIN   TYPE = "NOT IN"
	LIKE    TYPE = "LIKE"
	NOTLIKE TYPE = "NOTLIKE"
)

// Builder is responsible to dynamically create SQL query parts.
type Builder struct {
	Connection
}

// AndX Creates a conjunction of the given boolean expressions.
// Example:
//
//    // (u.type = ?) AND (u.role = ?)
//    expr.andX('u.type = ?', 'u.role = ?'))
func (Builder) AndX(parts ...string) Composite {
	return Composite{AND, parts}
}

//
//    OrX Creates a disjunction of the given boolean expressions.
//
//    Example:
//
//
//        // (u.type = ?) OR (u.role = ?)
//        qb.where(qb.expr().orX('u.type = ?', 'u.role = ?'))
//
//    @param mixed x Optional clause. Defaults = null, but requires
//                    at least one defined when converting to string.
//
//    @return \Doctrine\DBAL\Query\Expression\CompositeExpression
//
func (Builder) OrX(parts ...string) Composite {
	return Composite{OR, parts}
}

//    /**
//     * Creates a comparison expression.
//     *
//     * @param mixed  x        The left expression.
//     * @param string operator One of the ExpressionBuilder::* constants.
//     * @param mixed  y        The right expression.
//     *
//     * @return string
//     */
func (Builder) Comparison(x string, operator TYPE, y string) string {
	return x + " " + string(operator) + " " + y
}

//    /**
//     * Creates an equality comparison expression with the given arguments.
//     *
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> = <right expr>. Example:
//     *
//     *
//     *     // u.id = ?
//     *     expr.eq('u.id', '?')
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (Builder) Eq(x, y string) string {
	return Builder.Comparison(Builder{}, x, EQ, y)
}

//    /**
//     * Creates a non equality comparison expression with the given arguments.
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> <> <right expr>. Example:
//     *
//     *
//     *     // u.id <> 1
//     *     q.where(q.expr().neq('u.id', '1'))
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (b Builder) Neq(x, y string) string {
	return b.Comparison(x, NEQ, y)
}

//    /**
//     * Creates a lower-than comparison expression with the given arguments.
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> < <right expr>. Example:
//     *
//     *
//     *     // u.id < ?
//     *     q.where(q.expr().lt('u.id', '?'))
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (b Builder) Lt(x, y string) string {
	return b.Comparison(x, LT, y)
}

//    /**
//     * Creates a lower-than-equal comparison expression with the given arguments.
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> <= <right expr>. Example:
//     *
//     *
//     *     // u.id <= ?
//     *     q.where(q.expr().lte('u.id', '?'))
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (b Builder) Lte(x, y string) string {
	return b.Comparison(x, LTE, y)
}

//    /**
//     * Creates a greater-than comparison expression with the given arguments.
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> > <right expr>. Example:
//     *
//     *
//     *     // u.id > ?
//     *     q.where(q.expr().gt('u.id', '?'))
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (b Builder) Gt(x, y string) string {
	return b.Comparison(x, GT, y)
}

//    /**
//     * Creates a greater-than-equal comparison expression with the given arguments.
//     * First argument is considered the left expression and the second is the right expression.
//     * When converted to string, it will generated a <left expr> >= <right expr>. Example:
//     *
//     *
//     *     // u.id >= ?
//     *     q.where(q.expr().gte('u.id', '?'))
//     *
//     * @param mixed x The left expression.
//     * @param mixed y The right expression.
//     *
//     * @return string
//     */
func (b Builder) Gte(x, y string) string {
	return b.Comparison(x, GTE, y)
}

//    /**
//     * Creates an IS NULL expression with the given arguments.
//     *
//     * @param string x The field in string format to be restricted by IS NULL.
//     *
//     * @return string
//     */
func (b Builder) IsNull(x string) string {
	return x + " IS NULL"
}

//    /**
//     * Creates an IS NOT NULL expression with the given arguments.
//     *
//     * @param string x The field in string format to be restricted by IS NOT NULL.
//     *
//     * @return string
//     */
func (b Builder) IsNotNull(x string) string {
	return x + " IS NOT NULL"
}

//    /**
//     * Creates a LIKE() comparison expression with the given arguments.
//     *
//     * @param string x Field in string format to be inspected by LIKE() comparison.
//     * @param mixed  y Argument to be used in LIKE() comparison.
//     *
//     * @return string
//     */
func (b Builder) Like(x, y string) string {
	return b.Comparison(x, LIKE, y)
}

//    /**
//     * Creates a NOT LIKE() comparison expression with the given arguments.
//     *
//     * @param string x Field in string format to be inspected by NOT LIKE() comparison.
//     * @param mixed  y Argument to be used in NOT LIKE() comparison.
//     *
//     * @return string
//     */
func (b Builder) NotLike(x, y string) string {
	return b.Comparison(x, NOTLIKE, y)
}

//    /**
//     * Creates a IN () comparison expression with the given arguments.
//     *
//     * @param string       x The field in string format to be inspected by IN() comparison.
//     * @param string|array y The placeholder or the array of values to be used by IN() comparison.
//     *
//     * @return string
//     */
func (b Builder) In(x string, y ...interface{}) string {
	return b.Comparison(x, IN, "("+implode(", ", y...)+")")
}

//    /**
//     * Creates a NOT IN () comparison expression with the given arguments.
//     *
//     * @param string       x The field in string format to be inspected by NOT IN() comparison.
//     * @param string|array y The placeholder or the array of values to be used by NOT IN() comparison.
//     *
//     * @return string
//     */
func (b Builder) NotIn(x string, y ...interface{}) string {
	return b.Comparison(x, NOTIN, "("+implode(", ", y...)+")")
}

//    Literal Quotes a given input parameter.
func (b Builder) Literal(input interface{}, Type string) string {
	return b.Connection.Quote(input, Type)
}

func implode(separator string, values ...interface{}) string {
	result := []interface{}{}
	for i, value := range values {
		result = append(result, value)
		if i == (len(values) - 1) {
			break
		}
		result = append(result, separator)
	}
	return fmt.Sprint(result...)
}
