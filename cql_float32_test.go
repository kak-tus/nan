package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doCQLTest(t, gocql.TypeFloat, &v1, &v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doCQLTest(t, gocql.TypeFloat, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doCQLTest(t, gocql.TypeFloat, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}
