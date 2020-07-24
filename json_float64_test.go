package nan

import "testing"

func TestJSONNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doJSONTest(t, v1, v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}
