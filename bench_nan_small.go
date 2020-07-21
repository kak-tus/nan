package nan

//go:generate easyjson bench_nan_small.go

//easyjson:json
type nanSmall struct {
	Field000 NullString
	Field001 NullString
	Field002 NullString
	Field003 NullString
	Field004 NullString
	Field005 NullString
}

// Use separate value for encoding with jsoniter to ignore generated easyjson
// MarshalJSON and UnmarshalJSON
type nanSmallJSON struct {
	Field000 NullString
	Field001 NullString
	Field002 NullString
	Field003 NullString
	Field004 NullString
	Field005 NullString
}
