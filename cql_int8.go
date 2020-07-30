package nan

import "github.com/gocql/gocql"

func (n NullInt8) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return []byte{byte(n.Int8)}, nil
}

func (n *NullInt8) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 1 {
		*n = NullInt8{}
		return nil
	}

	dec := decTiny(data)

	*n = NullInt8{Valid: true, Int8: dec}

	return nil
}
