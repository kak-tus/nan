package nan

import "testing"

func TestSQLNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doSQLTest(t, v1, v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doSQLTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doSQLTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})
}
