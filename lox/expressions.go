package lox

type Binary struct {
	left     any
	operator Token
	right    any
}

type Grouping struct {
	expression any
}

type Literal struct {
	value any
}

type Unary struct {
	operator Token
	right    any
}
