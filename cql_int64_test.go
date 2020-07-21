package nan

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

func TestNullInt64(t *testing.T) {
	n1 := NullInt64{Valid: true, Int64: 7676}

	b, err := n1.MarshalCQL(gocql.TupleTypeInfo{})
	assert.NoError(t, err)

	var n2 NullInt64

	err = n2.UnmarshalCQL(gocql.TupleTypeInfo{}, b)
	assert.NoError(t, err)

	assert.Equal(t, n1, n2)
}
