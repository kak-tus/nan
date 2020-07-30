package nan

import "github.com/gocql/gocql"

func (n NullInt16) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return encShort(n.Int16), nil
}

func (n *NullInt16) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 2 {
		*n = NullInt16{}
		return nil
	}

	dec := decShort(data)

	*n = NullInt16{Valid: true, Int16: dec}

	return nil
}
