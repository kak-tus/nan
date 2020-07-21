package nan

import (
	"testing"

	"github.com/gocql/gocql"
)

func TestCQLNullBool(t *testing.T) {
	v1, v2 := true, true
	doCQLTest(t, gocql.TypeBoolean, &v1, &v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doCQLTest(t, gocql.TypeBoolean, &v1, &v2, &NullBool{Bool: false, Valid: true}, &NullBool{})

	doCQLTest(t, gocql.TypeBoolean, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doCQLTest(t, gocql.TypeBoolean, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})
}
