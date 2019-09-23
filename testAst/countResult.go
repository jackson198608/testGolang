package main

// 一个典型的中序遍历求解算法
func ExprASTResult(expr ExprAST) float64 {
	// 左右值
	var l, r float64
	switch expr.(type) {
	// 传入的根节点是 BinaryExprAST
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		switch ast.Lhs.(type) {
		case BinaryExprAST:
			// 递归左节点
			l = ExprASTResult(ast.Lhs)
		default:
			l = ast.Lhs.(NumberExprAST).Val
		}
		switch ast.Rhs.(type) {
		case BinaryExprAST:
			// 递归右节点
			r = ExprASTResult(ast.Rhs)
		default:
			r = ast.Rhs.(NumberExprAST).Val
		}
		// 现在 l,r 都有具体的值了，可以根据运算符运算
		switch ast.Op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		case "%":
			return float64(int(l) % int(r))
		default:

		}
	// 传入的根节点是 NumberExprAST,无需做任何事情，直接返回 Val 值
	case NumberExprAST:
		return expr.(NumberExprAST).Val
	}

	return 0.0
}
