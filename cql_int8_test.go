package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doCQLTest(t, gocql.TypeTinyInt, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doCQLTest(t, gocql.TypeTinyInt, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}
