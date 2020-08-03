package nan

import "testing"

func TestSQLNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doSQLTest(t, v1, v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doSQLTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doSQLTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}
