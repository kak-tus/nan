package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doCQLTest(t, gocql.TypeVarchar, &v1, &v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doCQLTest(t, gocql.TypeVarchar, &v1, &v2, &NullString{String: v1, Valid: true}, &NullString{})

	doCQLTest(t, gocql.TypeVarchar, nil, nil, &NullString{Valid: false}, &NullString{})
	doCQLTest(t, gocql.TypeVarchar, nil, nil, &NullString{Valid: false}, &NullString{Valid: true})
}
