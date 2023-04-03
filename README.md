# nan - **N**o **A**llocations **N**evermore

Package nan - Zero allocation Nullable structures in one library with handy conversion functions,
marshallers and unmarshallers.

[![Godoc](https://godoc.org/github.com/kak-tus/nan?status.svg)](https://pkg.go.dev/github.com/kak-tus/nan?tab=doc)
[![Coverage Status](https://coveralls.io/repos/github/kak-tus/nan/badge.svg?branch=master)](https://coveralls.io/github/kak-tus/nan?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kak-tus/nan)](https://goreportcard.com/report/github.com/kak-tus/nan)
![Go](https://github.com/kak-tus/nan/workflows/Go/badge.svg)

Features:
- short name "nan"
- handy conversion functions
- select which marshalers you want and limit dependencies to only those you actually need
- ability to convert your custom structs to nan compatible type with Valid field and all requested encoders/decoders

Supported types:
- bool
- float32
- float64
- int
- int8
- int16
- int32
- int64
- string
- time.Time
- uint
- uint8
- uint16
- uint32
- uint64
- more types will be added as necessary

Supported marshallers:
- Standart JSON
- encoding.TextMarshaler/TextUnmarshaler. Reuses standard JSON logic and format
- jsoniter
- easyjson
- go-json
- Scylla and Cassandra. Compatible with gocql
- SQL

## Usage

Simply create struct field or variable with one of the exported types and use it without any changes to external API.

JSON input/output will be converted to null or non null values. Scylla and Cassandra will use wire format compatible
with gocql.

```
	var data struct {
		Code nan.NullString `json:"code"`
	}

	b, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":null}
	println(string(b))

	data.Code = nan.String("1")
	// Equals to
	// data.Code = nan.NullString{String: "1", Valid: true}

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"1"}
	println(string(b))

	code := "2"

	// From addr. Can has value or be nil
	data.Code = nan.StringAddr(&code)

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"2"}
	println(string(b))

	// To usual value from nan
	codeVal := data.Code.String

	// 2
	println(codeVal)

	// To value addr from nan
	codeAddr := data.Code.Addr()

	// 2
	println(*codeAddr)
```

## Generate marshalers

```
# go install github.com/kak-tus/nan/cmd/nan@latest
# nan -help
```

Instead of depending on the whole github.com/kak-tus/nan you can also use `nan` command to select which marshalers you want. Simply run `nan` with one or more arguments and it will generate implementations for the specified marshalers in the current directory. For example, running
```
# nan gen -json -jsoniter
```
will generate nan.go, json.go, jsoniter.go files in the current working directory that contain only encoding/json and jsoniter marshalers. Nothing else will be generated so you don't have to depend on all the marshalers that github.com/kak-tus/nan supports. Generated files will use current directory name as its package name. You can also specify your own package name with `-pkg` argument.

## Custom structs generator

Imagine, that you have custom struct

```
type MyStruct struct {
	ID   int
	Name string
}
```

Use nan command on its file

```
# nan extra -json -jsoniter example/structs.go
```

This will generate *_nan.go near source files with json (or any other supported marshallers). And now you have nan compatible struct with all needed marshallers

```
var val MyStruct

nullVal := NanMyStruct(val)
// Equals to
// nullVal := NullMyStruct{MyStruct: val, Valid: true}

fmt.Println(nullVal.ID)
```

See [example](./example/README.md) to specific of easyjson, cql, sql generation.

nan extra coommand supports any number of file names at command line.

## Benchmarks

[See here](./bench/README.md).

## sqlc integration

[See here](./sqlc.md).
