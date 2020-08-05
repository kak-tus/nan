package nan

import "time"

//go:generate pkger -o cmd/nan

// NullInt64 - nullable int64
type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

// NullInt32 - nullable int32
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// NullTime - nullable time.Time
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// NullString - nullable string
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

// NullBool - nullable bool
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

// NullFloat64 - nullable float64
type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

// NullFloat32 - nullable float32
type NullFloat32 struct {
	Float32 float32
	Valid   bool // Valid is true if Float32 is not NULL
}

// NullInt16 - nullable int16
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}

// NullInt8 - nullable int8
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if Int8 is not NULL
}

// NullInt - nullable int
type NullInt struct {
	Int   int
	Valid bool // Valid is true if Int8 is not NULL
}
