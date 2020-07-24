package nan

import "github.com/kak-tus/nan"

//go:generate easyjson bench_nan_full.go

//easyjson:json
type nanFull struct {
	Field000 nan.NullString
	Field001 nan.NullBool
	Field002 nan.NullInt64
	Field003 nan.NullTime
	Field004 nan.NullFloat64
}

// Use separate value for encoding with jsoniter to ignore generated easyjson
// MarshalJSON and UnmarshalJSON
type nanFullJSON struct {
	Field000 nan.NullString
	Field001 nan.NullBool
	Field002 nan.NullInt64
	Field003 nan.NullTime
	Field004 nan.NullFloat64
}
