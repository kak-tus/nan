package nan

import "testing"

func TestSQLNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doSQLTest(t, v1, v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doSQLTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doSQLTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}
