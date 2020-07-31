package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	v1, v2 = int(76), int(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}
