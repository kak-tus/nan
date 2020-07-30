package nan

import (
	"math"

	"github.com/gocql/gocql"
)

func (n NullInt16) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int16))
}

func (n *NullInt16) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt16{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}
	if dec > math.MaxInt16 || dec < math.MinInt16 {
		return cqlMarshalErrorf("unmarshal int16: value %d out of range", dec)
	}
	*n = NullInt16{Valid: true, Int16: int16(dec)}

	return nil
}
