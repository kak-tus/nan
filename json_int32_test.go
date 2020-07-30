package nan

import "testing"

func TestJSONNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doJSONTest(t, v1, v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}
