package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
	importsTemplate  = "// imports"
	jsonTemplate     = "// JSON template"
	easyjsonTemplate = "// EasyJSON template"
	registerTemplate = "nan.nullTemplateType"
	helperTemplate   = "naninitialTemplateType"
)

type definition struct {
	name string
}

type visitor struct {
	definitions []definition
	functions   map[string]map[string]bool
	name        string
}

func generateExtra() {
	json := flag.Bool("json", false, "emit implementation of json.Marshaler and json.Unmarshaler")
	jsoniter := flag.Bool("jsoniter", false, "emit json-iterator encoder/decoder registration code")
	easyjson := flag.Bool("easyjson", false, "emit implementation of easyjson.Marshaler and easyjson.Unmarshaler")
	sql := flag.Bool("sql", false, "emit implementation of sql.Scanner and value")
	cql := flag.Bool("cql", false, "emit implementation of gocql.Marshaler and gocql.Unmarshaler")
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
	templateFilename = pkger.Include("/easyjson_template.go")
	dataEasyjson := readTemplate(templateFilename)
	templateFilename = pkger.Include("/sql_template.go")
	dataSQL := readTemplate(templateFilename)
	templateFilename = pkger.Include("/cql_template.go")
	dataCQL := readTemplate(templateFilename)

	files := flag.Args()

	if len(files) == 0 {
		return
	}

	currVisitor := visitor{
		definitions: make([]definition, 0),
		functions:   make(map[string]map[string]bool),
	}

	for _, file := range files {
		fset := token.NewFileSet()

		parsed, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			panic(err)
		}

		ast.Walk(&currVisitor, parsed)
	}

	if len(currVisitor.definitions) == 0 {
		return
	}

	out := fmt.Sprintf(genStr, strings.Join(os.Args[1:], " "))

	out += fmt.Sprintf("\npackage %s\n", *pkgName)

	// To add imports later
	out += "\n" + importsTemplate + "\n"

	imports := make([]string, 0)

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

		nameCall := "n." + def.name
		nameStruct := def.name + ":"
		helperTemplateTarget := "Nan" + def.name
		registerTemplateTarget := fmt.Sprintf("%s.%s", *pkgName, nameNan)

		replacer := strings.NewReplacer(
			initialTypeName, def.name,
			generateTypeName, nameNan,
			valueFieldName, def.name,
			valueCallName, nameCall,
			valueStructName, nameStruct,
			packageName, "",
			helperTemplate, helperTemplateTarget,
			registerTemplate, registerTemplateTarget,
		)

		withoutImport, imp := findImports(replacer.Replace(string(dataNan)))

		out += "\n" + withoutImport

		imports = append(imports, imp...)

		if *json {
			withoutImport, imp := findImports(replacer.Replace(string(dataJSON)))

			out += "\n" + withoutImport

			imports = append(imports, imp...)
		}

		if *jsoniter {
			withoutImport, imp := findImports(replacer.Replace(string(dataJsoniter)))

			out += "\n" + withoutImport

			imports = append(imports, imp...)
		}

		if *easyjson {
			withoutImport, imp := findImports(replacer.Replace(string(dataEasyjson)))

			switch {
			case currVisitor.hasFunction(def.name, "MarshalEasyJSON") && currVisitor.hasFunction(def.name, "UnmarshalEasyJSON"):
				withoutImport = removeLinesWithSuffix(withoutImport, jsonTemplate)
			case currVisitor.hasFunction(def.name, "MarshalJSON") && currVisitor.hasFunction(def.name, "UnmarshalJSON"):
				withoutImport = removeLinesWithSuffix(withoutImport, easyjsonTemplate)
			default:
				panic(fmt.Sprintf("Not found\n"+
					"MarshalEasyJSON/UnmarshalEasyJSON (preferred)\n"+
					"or MarshalJSON/UnmarshalJSON functions for %s\n"+
					"Run easyjson for your structs before run nan",
					def.name))
			}

			out += "\n" + withoutImport

			imports = append(imports, imp...)
		}

		if *sql {
			if !(currVisitor.hasFunction(def.name, "Value") && currVisitor.hasFunction(def.name, "Scan")) {
				panic(fmt.Sprintf("Not found\n"+
					"Value/Scan functions for %s\n"+
					"Write them for your structs before run nan",
					def.name))
			}

			withoutImport, imp := findImports(replacer.Replace(string(dataSQL)))

			out += "\n" + withoutImport

			imports = append(imports, imp...)
		}

		if *cql {
			if !(currVisitor.hasFunction(def.name, "MarshalCQL") && currVisitor.hasFunction(def.name, "UnmarshalCQL")) {
				panic(fmt.Sprintf("Not found\n"+
					"MarshalCQL/UnmarshalCQL functions for %s\n"+
					"Write them for your structs before run nan",
					def.name))
			}

			withoutImport, imp := findImports(replacer.Replace(string(dataCQL)))

			out += "\n" + withoutImport

			imports = append(imports, imp...)
		}
	}

	if len(imports) > 0 {
		sort.Strings(imports)

		joined := strings.Join(uniq(imports), "\n")
		formatted := fmt.Sprintf("import (\n%s\n)\n", joined)

		out = strings.Replace(out, importsTemplate, formatted, 1)
	}

	formatted, err := format.Source([]byte(out))
	if err != nil {
		panic(err)
	}

	genDir := filepath.Dir(files[0])

	resName := filepath.Join(genDir, *pkgName+"_nan.go")

	if err := ioutil.WriteFile(resName, formatted, 0644); err != nil {
		panic(err)
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
	case *ast.FuncDecl:
		if val.Recv == nil {
			return nil
		}

		for _, receiver := range val.Recv.List {
			var rname string

			switch rtype := receiver.Type.(type) {
			case *ast.Ident:
				rname = rtype.Name
			case *ast.StarExpr:
				rval, ok := rtype.X.(*ast.Ident)
				if ok {
					rname = rval.Name
				}
			}

			if rname == "" {
				continue
			}

			if _, ok := v.functions[rname]; !ok {
				v.functions[rname] = make(map[string]bool)
			}

			v.functions[rname][val.Name.Name] = true
		}

		return nil
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

func findImports(data string) (string, []string) {
	res, imports := findMultilineImports(data)
	if len(imports) > 0 {
		return res, imports
	}

	return findSingleImports(data)
}

func findMultilineImports(data string) (string, []string) {
	from := strings.Index(data, "\nimport (")
	if from < 0 {
		return data, nil
	}

	to := strings.Index(data[from:], "\n)")
	if to < 0 {
		return data, nil
	}

	imports := strings.Split(data[from+1:from+to+1], "\n")

	if len(imports) < 3 {
		return data, nil
	}

	withoutIncludes := data[:from] + "\n" + data[from+to+2:]

	// exclude first and last line from imports with "import (" and ")"
	return withoutIncludes, imports[1 : len(imports)-1]
}

func findSingleImports(data string) (string, []string) {
	from := strings.Index(data, "\nimport \"")
	if from < 0 {
		return data, nil
	}

	to := strings.Index(data[from:], "\"\n")
	if to < 0 {
		return data, nil
	}

	imports := data[from+7 : from+to+1]

	withoutIncludes := data[:from] + "\n" + data[from+to+2:]

	return withoutIncludes, []string{imports}
}

func removeLinesWithSuffix(s, suffix string) string {
	splitted := strings.Split(s, "\n")

	res := make([]string, 0, len(splitted))

	for _, line := range splitted {
		if strings.HasSuffix(line, suffix) {
			continue
		}

		res = append(res, line)
	}

	return strings.Join(res, "\n")
}

func (v *visitor) hasFunction(varType, function string) bool {
	if _, ok := v.functions[varType]; !ok {
		return false
	}

	if _, ok := v.functions[varType][function]; !ok {
		return false
	}

	return true
}

func uniq(strs []string) []string {
	vals := make(map[string]bool)

	for _, str := range strs {
		if str == "" {
			continue
		}

		vals[str] = true
	}

	res := make([]string, 0, len(vals))

	for k := range vals {
		res = append(res, k)
	}

	return res
}
