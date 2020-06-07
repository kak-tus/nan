package nan

func StringToNullString(v string) NullString {
	return NullString{String: v, Valid: true}
}
