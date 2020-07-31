package nan

import (
	"math"

	"github.com/gocql/gocql"
)

func (n NullInt32) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int32))
}

func (n *NullInt32) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt32{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}
	if dec > math.MaxInt32 || dec < math.MinInt32 {
		return cqlMarshalErrorf("unmarshal int32: value %d out of range", dec)
	}
	*n = NullInt32{Valid: true, Int32: int32(dec)}

	return nil
}
