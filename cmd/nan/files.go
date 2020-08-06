package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func generateForFiles(files []string) {
	for _, file := range files {
		fset := token.NewFileSet()

		parsed, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			panic(err)
		}

		ast.Inspect(parsed, func(node ast.Node) bool {
			switch val := node.(type) {
			case *ast.GenDecl:
				return true
			case *ast.Ident:
				println(val.Name)
				return true
			default:
				return true
			}
		})
	}
}
