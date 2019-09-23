package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 定义操作符优先级，value 越高，优先级越高
var precedence = map[string]int{"+": 20, "-": 20, "*": 40, "/": 40, "%": 40, "^": 60}

// AST 生成器结构体
type AST struct {
	// 词法分析的结果
	Tokens []*Token
	// 源字符串
	source string
	// 当前分析器分析的 Token
	currTok *Token
	// 当前分析器的位置
	currIndex int
	// 错误收集
	Err error
}

// 生成一个 AST 结构指针
func NewAST(toks []*Token, s string) *AST {
	a := &AST{
		Tokens: toks,
		source: s,
	}
	if a.Tokens == nil || len(a.Tokens) == 0 {
		a.Err = errors.New("empty token")
	} else {
		a.currIndex = 0
		a.currTok = a.Tokens[0]
	}
	return a
}

// 语法分析器入口
func (a *AST) ParseExpression() ExprAST {
	lhs := a.parsePrimary()
	return a.parseBinOpRHS(0, lhs)
}

// 获取下一个 Token
func (a *AST) getNextToken() *Token {
	a.currIndex++
	if a.currIndex < len(a.Tokens) {
		a.currTok = a.Tokens[a.currIndex]
		return a.currTok
	}
	return nil
}

// 获取操作优先级
func (a *AST) getTokPrecedence() int {
	if p, ok := precedence[a.currTok.Tok]; ok {
		return p
	}
	return -1
}

// 解析数字，并生成一个 NumberExprAST 节点
func (a *AST) parseNumber() NumberExprAST {
	f64, err := strconv.ParseFloat(a.currTok.Tok, 64)
	if err != nil {
		a.Err = fmt.Errorf("%v\nwant '(' or '0-9' but get '%s'\n%s",
			err.Error(),
			a.currTok.Tok,
			ErrPos(a.source, a.currTok.Offset))
		return NumberExprAST{}
	}
	n := NumberExprAST{
		Val: f64,
	}
	a.getNextToken()
	return n
}

// 获取一个节点，返回 ExprAST
// 这里会处理所有可能出现的类型，并对相应类型做解析
func (a *AST) parsePrimary() ExprAST {
	switch a.currTok.Type {
	case Literal:
		return a.parseNumber()
	case Operator:
		// 对 () 语法处理
		if a.currTok.Tok == "(" {
			a.getNextToken()
			e := a.ParseExpression()
			if e == nil {
				return nil
			}
			if a.currTok.Tok != ")" {
				a.Err = fmt.Errorf("want ')' but get %s\n%s",
					a.currTok.Tok,
					ErrPos(a.source, a.currTok.Offset))
				return nil
			}
			a.getNextToken()
			return e
		} else {
			return a.parseNumber()
		}
	default:
		return nil
	}
}

// 循环获取操作符的优先级，将高优先级的递归成较深的节点
// 这是生成正确的 AST 结构最重要的一个算法，一定要仔细阅读、理解
func (a *AST) parseBinOpRHS(execPrec int, lhs ExprAST) ExprAST {
	for {
		tokPrec := a.getTokPrecedence()
		if tokPrec < execPrec {
			return lhs
		}
		binOp := a.currTok.Tok
		if a.getNextToken() == nil {
			return lhs
		}
		rhs := a.parsePrimary()
		if rhs == nil {
			return nil
		}
		nextPrec := a.getTokPrecedence()
		if tokPrec < nextPrec {
			// 递归，将当前优先级+1
			rhs = a.parseBinOpRHS(tokPrec+1, rhs)
			if rhs == nil {
				return nil
			}
		}
		lhs = BinaryExprAST{
			Op:  binOp,
			Lhs: lhs,
			Rhs: rhs,
		}
	}
}
