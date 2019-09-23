package main

import (
	"fmt"
	"strconv"
)

// 基础表达式节点接口
type ExprAST interface {
	toStr() string
}

// 数字表达式节点
type NumberExprAST struct {
	// 具体的值
	Val float64
}

// 操作表达式节点
type BinaryExprAST struct {
	// 操作符
	Op string
	// 左右节点，可能是 数字表达式节点/操作表达式节点/nil
	Lhs,
	Rhs ExprAST
}

// 实现接口
func (n NumberExprAST) toStr() string {
	return fmt.Sprintf(
		"NumberExprAST:%s",
		strconv.FormatFloat(n.Val, 'f', 0, 64),
	)
}

// 实现接口
func (b BinaryExprAST) toStr() string {
	return fmt.Sprintf(
		"BinaryExprAST: (%s %s %s)",
		b.Op,
		b.Lhs.toStr(),
		b.Rhs.toStr(),
	)
}
