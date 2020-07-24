package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doCQLTest(t, gocql.TypeDouble, &v1, &v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doCQLTest(t, gocql.TypeDouble, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doCQLTest(t, gocql.TypeDouble, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}
