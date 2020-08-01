package nan

import (
	"testing"
	"time"
)

func TestSQLNullTime(t *testing.T) {
	v1, v2 := time.Now().UTC(), time.Now().UTC()
	doSQLTest(t, v1, v2, &NullTime{Time: v1, Valid: true}, &NullTime{})

	v1, v2 = time.Time{}, time.Time{}
	doSQLTest(t, v1, v2, &NullTime{Valid: true}, &NullTime{})

	doSQLTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doSQLTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})
}
