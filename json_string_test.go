package nan

import "testing"

func TestJSONNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	doJSONTest(t, nil, nil, &NullString{Valid: false}, &NullString{})
}
