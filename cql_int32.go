package nan

import "github.com/gocql/gocql"

func (n NullInt32) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return encInt(n.Int32), nil
}

func (n *NullInt32) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 4 {
		*n = NullInt32{}
		return nil
	}

	dec := decInt(data)

	*n = NullInt32{Valid: true, Int32: dec}

	return nil
}
