package nan

import (
	"math"

	"github.com/gocql/gocql"
)

func (n NullFloat64) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1025
	return encBigInt(int64(math.Float64bits(n.Float64))), nil
}

func (n *NullFloat64) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 8 {
		*n = NullFloat64{}
		return nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1043
	dec := math.Float64frombits(uint64(decBigInt(data)))

	*n = NullFloat64{Valid: true, Float64: dec}

	return nil
}
