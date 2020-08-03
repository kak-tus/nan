package nan

import "testing"

func TestSQLNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doSQLTest(t, v1, v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doSQLTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doSQLTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}
