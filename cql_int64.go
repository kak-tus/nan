package nan

import "github.com/gocql/gocql"

func (n NullInt64) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return encBigInt(n.Int64), nil
}

func (n *NullInt64) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 8 {
		*n = NullInt64{}
		return nil
	}

	dec := decBigInt(data)

	*n = NullInt64{Valid: true, Int64: dec}

	return nil
}
