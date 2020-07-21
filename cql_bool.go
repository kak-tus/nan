package nan

import "github.com/gocql/gocql"

func (n NullBool) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	if n.Bool {
		return []byte{1}, nil
	}

	return []byte{0}, nil
}

func (n *NullBool) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) == 0 {
		*n = NullBool{}
		return nil
	}

	*n = NullBool{Valid: true, Bool: data[0] != 0}

	return nil
}
