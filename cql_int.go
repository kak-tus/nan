package nan

import (
	"math"

	"github.com/gocql/gocql"
)

func (n NullInt) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int))
}

func (n *NullInt) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}
	if ^uint(0) == math.MaxUint32 && (dec > math.MaxInt32 || dec < math.MinInt32) {
		return cqlMarshalErrorf("unmarshal int: value %d out of range", dec)
	}
	*n = NullInt{Valid: true, Int: int(dec)}

	return nil
}
