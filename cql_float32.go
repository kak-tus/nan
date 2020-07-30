package nan

import (
	"math"

	"github.com/gocql/gocql"
)

func (n NullFloat32) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L982
	return encInt(int32(math.Float32bits(n.Float32))), nil
}

func (n *NullFloat32) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 4 {
		*n = NullFloat32{}
		return nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1002
	dec := math.Float32frombits(uint32(decInt(data)))

	*n = NullFloat32{Valid: true, Float32: dec}

	return nil
}
