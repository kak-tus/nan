package nan

import (
	"testing"
)

func TestSQLNullBool(t *testing.T) {
	v1, v2 := true, true
	doSQLTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doSQLTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	doSQLTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doSQLTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})
}
