package nan

import "testing"

func TestJSONNullBool(t *testing.T) {
	v1, v2 := true, true
	doJSONTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doJSONTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	doJSONTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doJSONTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})
}
