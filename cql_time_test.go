package nan

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
)

func TestCQLNullTime(t *testing.T) {
	// Cassandra truncates it
	v1 := time.Now().UTC().Truncate(time.Millisecond)
	doCQLTest(t, gocql.TypeTimestamp, &v1, &time.Time{}, &NullTime{Time: v1, Valid: true}, &NullTime{})

	v1 = time.Time{}
	doCQLTest(t, gocql.TypeTimestamp, &v1, &time.Time{}, &NullTime{Valid: true}, &NullTime{})

	doCQLTest(t, gocql.TypeTimestamp, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doCQLTest(t, gocql.TypeTimestamp, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})
}
