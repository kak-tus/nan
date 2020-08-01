package nan

import "testing"

func TestSQLNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doSQLTest(t, v1, v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doSQLTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doSQLTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}
