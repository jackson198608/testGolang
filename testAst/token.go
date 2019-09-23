package main

const (
	// 字面量，e.g. 50
	Literal = iota
	// 操作符, e.g. + - * /
	Operator
)

type Token struct {
	// 原始字符
	Tok string
	// 类型，有 Literal、Operator 两种
	Type   int
	Offset int
}
