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
