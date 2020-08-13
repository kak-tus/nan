package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/markbates/pkger"
)

const (
	initialTypeName  = "initialTemplateType"
	generateTypeName = "nullTemplateType"
	valueFieldName   = "Value"
)

func generateForFiles(files []string) {
	for _, file := range files {
		fset := token.NewFileSet()

		parsed, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			panic(err)
		}

		filenameNan := pkger.Include("/nan_template.go")

		fileNan, err := pkger.Open(filenameNan)
		if err != nil {
			panic(err)
		}

		defer fileNan.Close()

		dataNan, err := ioutil.ReadAll(fileNan)
		if err != nil {
			panic(err)
		}

		var name string

		ast.Inspect(parsed, func(node ast.Node) bool {
			switch val := node.(type) {
			case *ast.File:
				// fmt.Println("file", val.Name.String())
				return true
			case *ast.GenDecl:
				// fmt.Println("decl", val.Tok.String())
				return true
			case *ast.TypeSpec:
				// fmt.Println("spec", val.Name.String())
				return true
			case *ast.Ident:
				// fmt.Println("ident", val.Name)
				name = val.Name
				return true
			case *ast.StructType:
				if name == "" {
					return false
				}

				var buf bytes.Buffer

				printer.Fprint(&buf, fset, val)
				println(buf.String())

				nameUC := strings.ToUpper(string([]rune(name)[0])) + string([]rune(name)[1:])
				nameNan := "Null" + nameUC

				replacer := strings.NewReplacer(initialTypeName, name, generateTypeName, nameNan, valueFieldName, nameUC)
				out := replacer.Replace(string(dataNan))

				fmt.Println(out)

				// fmt.Println("struct", val.Fields.NumFields())
				// printer.Fprint(os.Stdout, fset, node)
				return true
			default:
				// printer.Fprint(os.Stdout, fset, node)
				// fset.Position(val.Pos()).
				// fmt.Println(val)
				// return true
				return false
			}
		})
	}
}
