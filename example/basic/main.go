package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/kak-tus/nan"
)

func main() {
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
}
