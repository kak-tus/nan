package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kak-tus/nan"
)

func generateDefault() {
	cql := flag.Bool("cql", false, "emit implementation of gocql.Marshaler and gocql.Unmarshaler")
	easyjson := flag.Bool("easyjson", false, "emit implementation of easyjson.Marshaler and easyjson.Unmarshaler")
	json := flag.Bool("json", false, "emit implementation of json.Marshaler and json.Unmarshaler. Conflicts with -goccyjson")
	jsoniter := flag.Bool("jsoniter", false, "emit json-iterator encoder/decoder registration code")
	goccyjson := flag.Bool("goccyjson", false, "emit implementation of json.Marshaler and json.Unmarshaler that uses goccy/go-json. Conflicts with -json")
	sql := flag.Bool("sql", false, "emit implementation of sql.Scanner and value")
	text := flag.Bool("text", false, "emit implementation of encoding.TextMarshaler and encoding.TextUnmarshaler")
	pkgName := flag.String("pkg", "", "specify generated package name. By default will use working directory name")

	flag.Parse()

	if *json && *goccyjson {
		fmt.Println("-json and -goccyjson are mutually exclusive. Choose one")
		return
	}

	if *pkgName == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		_, *pkgName = filepath.Split(wd)
	}

	files := make([]string, 0)
	if *cql {
		files = append(files, "cql.go")
		files = append(files, "cql_helpers.go")
	}

	if *easyjson {
		files = append(files, "easyjson.go")
	}

	if *json {
		files = append(files, "json.go")
	}

	if *jsoniter {
		files = append(files, "jsoniter.go")
	}

	if *goccyjson {
		files = append(files, "json.go")
	}

	if *sql {
		files = append(files, "sql.go")
		files = append(files, "sql_convert.go")
	}

	if *text {
		files = append(files, "text.go")
	}

	if len(files) == 0 {
		flag.PrintDefaults()
		return
	}

	files = append(
		files,
		"nan.go",
		"helpers.go",
		"LICENSE",
	)

	for i := range files {
		in, err := nan.EmbeddedSources.Open(files[i])
		if err != nil {
			panic(err)
		}

		src := []byte(fmt.Sprintf(genStr, strings.Join(os.Args[1:], " ")))

		reader := bufio.NewReader(in)

		for {
			line, err := reader.ReadBytes('\n')
			if strings.HasPrefix(string(line), "package ") {
				line = []byte(fmt.Sprintf("package %s\n", *pkgName))
			} else if strings.HasPrefix(string(line), "//go:generate ") {
				continue
			}

			if *goccyjson {
				line = convertJsonToGoccyJson(line, true)
			}

			src = append(src, line...)

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}
		}

		if err := ioutil.WriteFile(files[i], src, 0644); err != nil {
			panic(err)
		}

		in.Close()
	}
}
