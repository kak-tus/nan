package example

type MyStruct2 struct {
	Val1 NullMyStruct
	Val2 MyStruct
	Val3 NullMyStruct
	Val4 MyStruct
}

// Struct for test jsoniter encoding
// Because easyjson generates marshallers for encode/decode and original
// jsoniter registered functions not used
//easyjson:skip
type MyStruct3 struct {
	Val1 NullMyStruct
	Val2 MyStruct
	Val3 NullMyStruct
	Val4 MyStruct
}
