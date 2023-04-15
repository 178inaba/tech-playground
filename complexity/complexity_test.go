package complexity

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestComplexity(t *testing.T) {
	testcases := []struct {
		name       string
		code       string
		complexity int
	}{}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			a := getAST(t, testcase.code)
			c := Count(a)
			if c != testcase.complexity {
				t.Errorf("got=%d, want=%d", c, testcase.complexity)
			}
		})
	}
}

func getAST(t *testing.T, code string) ast.Node {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, decl := range file.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatal("no function declear found")
	return nil
}
