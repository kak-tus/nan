package nan

import "github.com/kak-tus/nan"

//go:generate easyjson bench_nan_small.go

//easyjson:json
type nanSmall struct {
	Field000 nan.NullString
	Field001 nan.NullString
	Field002 nan.NullString
	Field003 nan.NullString
	Field004 nan.NullString
	Field005 nan.NullString
}

// Use separate value for encoding with jsoniter to ignore generated easyjson
// MarshalJSON and UnmarshalJSON
type nanSmallJSON struct {
	Field000 nan.NullString
	Field001 nan.NullString
	Field002 nan.NullString
	Field003 nan.NullString
	Field004 nan.NullString
	Field005 nan.NullString
}
