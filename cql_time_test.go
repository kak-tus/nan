package nan

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

func TestNullTime(t *testing.T) {
	// Cassandra truncates it
	tm := time.Now().UTC().Truncate(time.Millisecond)

	n1 := NullTime{Valid: true, Time: tm}

	b, err := n1.MarshalCQL(gocql.TupleTypeInfo{})
	assert.NoError(t, err)

	var n2 NullTime

	err = n2.UnmarshalCQL(gocql.TupleTypeInfo{}, b)
	assert.NoError(t, err)

	assert.Equal(t, n1, n2)
}
