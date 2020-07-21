package nan

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

func TestNullBool(t *testing.T) {
	nan1 := NullBool{Bool: true, Valid: true}

	encoded1, err := nan1.MarshalCQL(gocql.TupleTypeInfo{})
	assert.NoError(t, err)

	cType := gocql.NewNativeType(5, gocql.TypeBoolean, "")

	val1 := true

	encoded2, err := gocql.Marshal(cType, &val1)
	assert.NoError(t, err)

	assert.Equal(t, encoded1, encoded2)

	var nan2 NullBool

	err = nan2.UnmarshalCQL(gocql.TupleTypeInfo{}, encoded1)
	assert.NoError(t, err)

	assert.Equal(t, nan1, nan2)

	var val2 bool

	err = gocql.Unmarshal(cType, encoded2, &val2)
	assert.NoError(t, err)

	assert.Equal(t, val1, val2)

	assert.Equal(t, nan1.Bool, val2)
}
