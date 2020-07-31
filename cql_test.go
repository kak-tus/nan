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
	cTypeEnc := gocql.NewNativeType(5, cType, "")

	encodedNan, err := nanVal1.MarshalCQL(cTypeEnc)
	assert.NoError(t, err)

	encodedVal, err := gocql.Marshal(cTypeEnc, val1)
	assert.NoError(t, err)

	assert.Equal(t, encodedNan, encodedVal)

	err = nanVal2.UnmarshalCQL(cTypeEnc, encodedNan)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)

	// not decode nil values
	if val2 != nil {
		err = gocql.Unmarshal(cTypeEnc, encodedVal, val2)
		assert.NoError(t, err)

		assert.Equal(t, val1, val2)
	}
}
