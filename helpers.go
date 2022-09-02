package nan

import "time"

// StringToNullString - converts string to NullString
//
// Deprecated: use shorter variant
func StringToNullString(v string) NullString {
	return NullString{String: v, Valid: true}
}

// BoolToNullBool - converts bool to NullBool
//
// Deprecated: use shorter variant
func BoolToNullBool(v bool) NullBool {
	return NullBool{Bool: v, Valid: true}
}

// Int64ToNullInt64 - converts int64 to NullInt64
//
// Deprecated: use shorter variant
func Int64ToNullInt64(v int64) NullInt64 {
	return NullInt64{Int64: v, Valid: true}
}

// TimeToNullTime - converts time.Time to NullTime
//
// Deprecated: use shorter variant
func TimeToNullTime(v time.Time) NullTime {
	return NullTime{Time: v, Valid: true}
}

// Float64ToNullFloat64 - converts float64 to NullFloat64
//
// Deprecated: use shorter variant
func Float64ToNullFloat64(v float64) NullFloat64 {
	return NullFloat64{Float64: v, Valid: true}
}

// Float32ToNullFloat32 - converts float32 to NullFloat32
//
// Deprecated: use shorter variant
func Float32ToNullFloat32(v float32) NullFloat32 {
	return NullFloat32{Float32: v, Valid: true}
}

// Int8ToNullInt8 - converts int8 to NullInt8
//
// Deprecated: use shorter variant
func Int8ToNullInt8(v int8) NullInt8 {
	return NullInt8{Int8: v, Valid: true}
}

// Int16ToNullInt16 - converts int16 to NullInt16
//
// Deprecated: use shorter variant
func Int16ToNullInt16(v int16) NullInt16 {
	return NullInt16{Int16: v, Valid: true}
}

// Int32ToNullInt32 - converts int32 to NullInt32
//
// Deprecated: use shorter variant
func Int32ToNullInt32(v int32) NullInt32 {
	return NullInt32{Int32: v, Valid: true}
}

// IntToNullInt - converts int to NullInt
//
// Deprecated: use shorter variant
func IntToNullInt(v int) NullInt {
	return NullInt{Int: v, Valid: true}
}

// String - converts string to NullString
func String(v string) NullString {
	return NullString{String: v, Valid: true}
}

// Bool - converts bool to NullBool
func Bool(v bool) NullBool {
	return NullBool{Bool: v, Valid: true}
}

// Int64 - converts int64 to NullInt64
func Int64(v int64) NullInt64 {
	return NullInt64{Int64: v, Valid: true}
}

// Time - converts time.Time to NullTime
func Time(v time.Time) NullTime {
	return NullTime{Time: v, Valid: true}
}

// Float64 - converts float64 to NullFloat64
func Float64(v float64) NullFloat64 {
	return NullFloat64{Float64: v, Valid: true}
}

// Float32 - converts float32 to NullFloat32
func Float32(v float32) NullFloat32 {
	return NullFloat32{Float32: v, Valid: true}
}

// Int8 - converts int8 to NullInt8
func Int8(v int8) NullInt8 {
	return NullInt8{Int8: v, Valid: true}
}

// Int16 - converts int16 to NullInt16
func Int16(v int16) NullInt16 {
	return NullInt16{Int16: v, Valid: true}
}

// Int32 - converts int32 to NullInt32
func Int32(v int32) NullInt32 {
	return NullInt32{Int32: v, Valid: true}
}

// Int - converts int to NullInt
func Int(v int) NullInt {
	return NullInt{Int: v, Valid: true}
}

// StringAddr - converts string address to NullString
func StringAddr(v *string) NullString {
	if v == nil {
		return NullString{}
	}

	return NullString{String: *v, Valid: true}
}

// BoolAddr - converts bool address to NullBool
func BoolAddr(v *bool) NullBool {
	if v == nil {
		return NullBool{}
	}

	return NullBool{Bool: *v, Valid: true}
}

// Int64Addr - converts int64 address to NullInt64
func Int64Addr(v *int64) NullInt64 {
	if v == nil {
		return NullInt64{}
	}

	return NullInt64{Int64: *v, Valid: true}
}

// TimeAddr - converts time.Time address to NullTime
func TimeAddr(v *time.Time) NullTime {
	if v == nil {
		return NullTime{}
	}

	return NullTime{Time: *v, Valid: true}
}

// Float64Addr - converts float64 address to NullFloat64
func Float64Addr(v *float64) NullFloat64 {
	if v == nil {
		return NullFloat64{}
	}

	return NullFloat64{Float64: *v, Valid: true}
}

// Float32Addr - converts float32 address to NullFloat32
func Float32Addr(v *float32) NullFloat32 {
	if v == nil {
		return NullFloat32{}
	}

	return NullFloat32{Float32: *v, Valid: true}
}

// Int8Addr - converts int8 address to NullInt8
func Int8Addr(v *int8) NullInt8 {
	if v == nil {
		return NullInt8{}
	}

	return NullInt8{Int8: *v, Valid: true}
}

// Int16Addr - converts int16 address to NullInt16
func Int16Addr(v *int16) NullInt16 {
	if v == nil {
		return NullInt16{}
	}

	return NullInt16{Int16: *v, Valid: true}
}

// Int32Addr - converts int32 address to NullInt32
func Int32Addr(v *int32) NullInt32 {
	if v == nil {
		return NullInt32{}
	}

	return NullInt32{Int32: *v, Valid: true}
}

// IntAddr - converts int address to NullInt
func IntAddr(v *int) NullInt {
	if v == nil {
		return NullInt{}
	}

	return NullInt{Int: *v, Valid: true}
}
