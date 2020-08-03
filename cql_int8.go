package nan

import (
	"math"

	"github.com/gocql/gocql"
)

// MarshalCQL - marshaller for cql
func (n NullInt8) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int8))
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullInt8) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt8{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}
	if dec > math.MaxInt8 || dec < math.MinInt8 {
		return cqlMarshalErrorf("unmarshal int8: value %d out of range", dec)
	}
	*n = NullInt8{Valid: true, Int8: int8(dec)}

	return nil
}
