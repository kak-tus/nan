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
		return nil
	}

	dec := decBigInt(data)

	*n = NullInt64{Valid: true, Int64: dec}

	return nil
}

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L598
func encBigInt(x int64) []byte {
	return []byte{byte(x >> 56), byte(x >> 48), byte(x >> 40), byte(x >> 32),
		byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)}
}

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L908
func decBigInt(data []byte) int64 {
	if len(data) != 8 {
		return 0
	}

	return int64(data[0])<<56 | int64(data[1])<<48 |
		int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 |
		int64(data[6])<<8 | int64(data[7])
}
