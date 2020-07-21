package nan

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	n1 := NullString{Valid: true, String: "7676"}

	b, err := n1.MarshalCQL(gocql.TupleTypeInfo{})
	assert.NoError(t, err)

	var n2 NullString

	err = n2.UnmarshalCQL(gocql.TupleTypeInfo{}, b)
	assert.NoError(t, err)

	assert.Equal(t, n1, n2)
}
