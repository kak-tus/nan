package nan

import "testing"

func TestJSONNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doJSONTest(t, v1, v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}
