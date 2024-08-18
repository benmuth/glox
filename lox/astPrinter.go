package lox

import (
	"fmt"
	"strconv"
	"strings"
)

func parenthesize(name string, exprs ...Expression) string {
	var b strings.Builder

	fmt.Fprintf(&b, "(%s", name)
	for _, expr := range exprs {
		b.Write([]byte(fmt.Sprintf(" %s", expr.String())))
	}
	b.Write([]byte{')'})
	return b.String()
}

func (b Binary) String() string {
	return parenthesize(b.operator.Lexeme, b.left, b.right)
}

func (g Grouping) String() string {
	return parenthesize("group", g.expression)
}

func (l Literal) String() string {
	if l.value == nil {
		return "nil"
	}
	switch l.value.(type) {
	case string:
		return l.value.(string)
	case int:
		return strconv.Itoa(l.value.(int))
	case float64:
		return strconv.FormatFloat(l.value.(float64), 'f', 2, 64)
	default:
		panic("invalid literal")
	}

}

func (u Unary) String() string {
	return parenthesize(u.operator.Lexeme, u.right)
}
