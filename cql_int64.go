package nan

import (
	"github.com/gocql/gocql"
)

func (n NullInt64) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, n.Int64)
}

func (n *NullInt64) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt64{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}
	*n = NullInt64{Valid: true, Int64: dec}

	return nil
}
