package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	v1, v2 = int32(76), int32(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}
