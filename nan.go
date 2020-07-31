package nan

import "database/sql"

type NullInt64 sql.NullInt64
type NullInt32 sql.NullInt32
type NullTime sql.NullTime
type NullString sql.NullString
type NullBool sql.NullBool
type NullFloat64 sql.NullFloat64

type NullFloat32 struct {
	Float32 float32
	Valid   bool // Valid is true if Float32 is not NULL
}

type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

type NullInt struct {
	Int   int
	Valid bool // Valid is true if Int8 is not NULL
}
