package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	v1, v2 = int16(76), int16(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doCQLTest(t, gocql.TypeSmallInt, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doCQLTest(t, gocql.TypeSmallInt, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})
}
