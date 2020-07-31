package nan

import "time"

type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

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
