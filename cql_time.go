package nan

import (
	"time"

	"github.com/gocql/gocql"
)

func (n NullTime) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1160
	x := int64(n.Time.UTC().Unix()*1e3) + int64(n.Time.UTC().Nanosecond()/1e6)

	enc := encBigInt(x)

	return enc, nil
}

func (n *NullTime) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 8 {
		return nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1213
	dec := decBigInt(data)

	sec := dec / 1000
	nsec := (dec - sec*1000) * 1000000
	t := time.Unix(sec, nsec).In(time.UTC)

	*n = NullTime{Valid: true, Time: t}

	return nil
}
