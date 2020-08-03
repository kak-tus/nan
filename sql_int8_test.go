package nan

import "testing"

func TestSQLNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doSQLTest(t, v1, v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doSQLTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doSQLTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}
