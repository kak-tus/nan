package nan

import "time"

func StringToNullString(v string) NullString {
	return NullString{String: v, Valid: true}
}

func BoolToNullBool(v bool) NullBool {
	return NullBool{Bool: v, Valid: true}
}

func Int64ToNullInt64(v int64) NullInt64 {
	return NullInt64{Int64: v, Valid: true}
}

func TimeToNullTime(v time.Time) NullTime {
	return NullTime{Time: v, Valid: true}
}

func Float64ToNullFloat64(v float64) NullFloat64 {
	return NullFloat64{Float64: v, Valid: true}
}

func Float32ToNullFloat32(v float32) NullFloat32 {
	return NullFloat32{Float32: v, Valid: true}
}

func Int8ToNullInt8(v int8) NullInt8 {
	return NullInt8{Int8: v, Valid: true}
}

func Int16ToNullInt16(v int16) NullInt16 {
	return NullInt16{Int16: v, Valid: true}
}

func Int32ToNullInt32(v int32) NullInt32 {
	return NullInt32{Int32: v, Valid: true}
}
