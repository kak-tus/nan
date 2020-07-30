package nan

import "testing"

func TestJSONNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doJSONTest(t, v1, v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}
