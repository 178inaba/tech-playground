package complexity

import "go/ast"

func Count(node ast.Node) int {
	count := 1

	ast.Inspect(node, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.IfStmt:
			count++
		}
		return true
	})

	return count
}
