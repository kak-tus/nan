package nan

import (
	"testing"
	"time"
)

func TestJSONNullTime(t *testing.T) {
	v1 := time.Now().UTC()
	doJSONTest(t, v1, &time.Time{}, &NullTime{Time: v1, Valid: true}, &NullTime{})

	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})
}
