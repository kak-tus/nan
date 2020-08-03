package nan

import "testing"

func TestSQLNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doSQLTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doSQLTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	doSQLTest(t, nil, nil, &NullString{Valid: false}, &NullString{})
	doSQLTest(t, nil, nil, &NullString{Valid: false}, &NullString{Valid: true})
}
