package lox

type Binary struct {
	left     Expression
	operator Token
	right    Expression
}

type Grouping struct {
	expression Expression
}

type Literal struct {
	value any
}

type Unary struct {
	operator Token
	right    Expression
}

type Expression interface {
	String() string
}
