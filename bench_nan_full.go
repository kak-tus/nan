package nan

//go:generate easyjson bench_nan_full.go

//easyjson:json
type nanFull struct {
	Field000 NullString
	Field001 NullBool
	Field002 NullInt64
	Field003 NullTime
	Field004 NullFloat64
}

// Use separate value for encoding with jsoniter to ignore generated easyjson
// MarshalJSON and UnmarshalJSON
type nanFullJSON struct {
	Field000 NullString
	Field001 NullBool
	Field002 NullInt64
	Field003 NullTime
	Field004 NullFloat64
}
