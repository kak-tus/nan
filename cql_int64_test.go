package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	v1, v2 = int64(76), int64(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doCQLTest(t, gocql.TypeBigInt, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doCQLTest(t, gocql.TypeBigInt, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})
}
