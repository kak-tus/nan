package nan

//go:generate easyjson bench_pointer_small.go

//easyjson:json
type pointerSmall struct {
	Field000 *string
	Field001 *string
	Field002 *string
	Field003 *string
	Field004 *string
	Field005 *string
}

// Use separate value for encoding with jsoniter to ignore generated easyjson
// MarshalJSON and UnmarshalJSON
type pointerSmallJSON struct {
	Field000 *string
	Field001 *string
	Field002 *string
	Field003 *string
	Field004 *string
	Field005 *string
}
