package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger"
)

const (
	initialTypeName  = "initialTemplateType"
	generateTypeName = "nullTemplateType"
	valueFieldName   = "NullTemplateValue initialTemplateType"
	valueCallName    = "n.NullTemplateValue"
	valueStructName  = "NullTemplateValue:"
	packageName      = "package nan"
)

type definition struct {
	name string
}

type visitor struct {
	name        string
	definitions []definition
}

func generateExtra() {
	json := flag.Bool("json", false, "emit implementation of json.Marshaler and json.Unmarshaler")
	jsoniter := flag.Bool("jsoniter", false, "emit json-iterator encoder/decoder registration code")
	pkgName := flag.String("pkg", "", "specify generated package name. By default will use working directory name")

	flag.Parse()

	if *pkgName == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		_, *pkgName = filepath.Split(wd)
	}

	var templateFilename string

	templateFilename = pkger.Include("/nan_template.go")
	dataNan := readTemplate(templateFilename)
	templateFilename = pkger.Include("/json_template.go")
	dataJSON := readTemplate(templateFilename)
	templateFilename = pkger.Include("/jsoniter_template.go")
	dataJsoniter := readTemplate(templateFilename)

	files := flag.Args()

	for _, file := range files {
		fset := token.NewFileSet()

		parsed, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			panic(err)
		}

		currVisitor := visitor{
			definitions: []definition{},
		}

		ast.Walk(&currVisitor, parsed)

		if len(currVisitor.definitions) == 0 {
			continue
		}

		out := fmt.Sprintf(genStr, strings.Join(os.Args[1:], " "))

		for _, def := range currVisitor.definitions {
			var nullPrefix string

			// Name starts with uppercased letter - exported
			if strings.ToUpper(string([]rune(def.name)[0])) == string([]rune(def.name)[0]) {
				nullPrefix = "Null"
			} else {
				nullPrefix = "null"
			}

			nameUC := strings.ToUpper(string([]rune(def.name)[0])) + string([]rune(def.name)[1:])
			nameNan := nullPrefix + nameUC

			namePkg := "package " + *pkgName
			nameCall := "n." + def.name
			nameStruct := def.name + ":"

			replacer := strings.NewReplacer(
				initialTypeName, def.name,
				generateTypeName, nameNan,
				valueFieldName, def.name,
				valueCallName, nameCall,
				valueStructName, nameStruct,
				packageName, namePkg,
			)

			out += replacer.Replace(string(dataNan))

			// for other files replace "package" to empty string
			replacer = strings.NewReplacer(
				initialTypeName, def.name,
				generateTypeName, nameNan,
				valueFieldName, def.name,
				valueCallName, nameCall,
				valueStructName, nameStruct,
				packageName, "",
			)

			if *json {
				out += "\n" + replacer.Replace(string(dataJSON))
			}

			if *jsoniter {
				out += "\n" + replacer.Replace(string(dataJsoniter))
			}
		}

		ext := filepath.Ext(file)

		resName := file[0:len(file)-len(ext)] + "_nan" + ext

		if err := ioutil.WriteFile(resName, []byte(out), 0644); err != nil {
			panic(err)
		}
	}
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch val := node.(type) {
	case *ast.File:
		return v
	case *ast.GenDecl:
		return v
	case *ast.TypeSpec:
		return v
	case *ast.Ident:
		v.name = val.Name
		return v
	case *ast.StructType:
		if v.name == "" {
			return nil
		}

		v.definitions = append(v.definitions, definition{name: v.name})
		v.name = ""
	}

	return nil
}

func readTemplate(name string) []byte {
	templateFile, err := pkger.Open(name)
	if err != nil {
		panic(err)
	}

	defer templateFile.Close()

	data, err := ioutil.ReadAll(templateFile)
	if err != nil {
		panic(err)
	}

	return data
}
