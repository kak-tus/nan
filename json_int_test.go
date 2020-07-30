package nan

import "testing"

func TestJSONNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doJSONTest(t, v1, v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}
