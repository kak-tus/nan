package nan

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

type cqlNanCoder interface {
	MarshalCQL(gocql.TypeInfo) ([]byte, error)
	UnmarshalCQL(gocql.TypeInfo, []byte) error
}

func doCQLTest(t *testing.T, cType gocql.Type, val1, val2 interface{}, nanVal1, nanVal2 cqlNanCoder) {
	encodedNan, err := nanVal1.MarshalCQL(gocql.TupleTypeInfo{})
	assert.NoError(t, err)

	cTypeEnc := gocql.NewNativeType(5, cType, "")

	encodedVal, err := gocql.Marshal(cTypeEnc, val1)
	assert.NoError(t, err)

	assert.Equal(t, encodedNan, encodedVal)

	err = nanVal2.UnmarshalCQL(gocql.TupleTypeInfo{}, encodedNan)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)

	// not decode nil values
	if val1 != nil {
		err = gocql.Unmarshal(cTypeEnc, encodedVal, val2)
		assert.NoError(t, err)

		assert.Equal(t, val1, val2)
	}
}
