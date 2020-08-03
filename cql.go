package nan

import (
	"math"
	"time"

	"github.com/gocql/gocql"
)

// MarshalCQL - marshaller for cql
func (n NullBool) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	if n.Bool {
		return []byte{1}, nil
	}

	return []byte{0}, nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullBool) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) == 0 {
		*n = NullBool{}
		return nil
	}

	*n = NullBool{Valid: true, Bool: data[0] != 0}

	return nil
}

// MarshalCQL - marshaller for cql
func (n NullFloat32) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L982
	return encInt(int32(math.Float32bits(n.Float32))), nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullFloat32) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 4 {
		*n = NullFloat32{}
		return nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1002
	dec := math.Float32frombits(uint32(decInt(data)))

	*n = NullFloat32{Valid: true, Float32: dec}

	return nil
}

// MarshalCQL - marshaller for cql
func (n NullFloat64) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1025
	return encBigInt(int64(math.Float64bits(n.Float64))), nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullFloat64) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) != 8 {
		*n = NullFloat64{}
		return nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1043
	dec := math.Float64frombits(uint64(decBigInt(data)))

	*n = NullFloat64{Valid: true, Float64: dec}

	return nil
}

// MarshalCQL - marshaller for cql
func (n NullInt) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int))
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullInt) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}

	if ^uint(0) == math.MaxUint32 && (dec > math.MaxInt32 || dec < math.MinInt32) {
		return cqlMarshalErrorf("unmarshal int: value %d out of range", dec)
	}

	*n = NullInt{Valid: true, Int: int(dec)}

	return nil
}

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

// MarshalCQL - marshaller for cql
func (n NullInt16) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int16))
}

// UnmarshalCQL - unmarshaller for cql
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

// MarshalCQL - marshaller for cql
func (n NullInt32) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, int64(n.Int32))
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullInt32) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullInt32{}
		return nil
	}

	dec, err := unmarshalIntLike(info, data)
	if err != nil {
		return err
	}

	if dec > math.MaxInt32 || dec < math.MinInt32 {
		return cqlMarshalErrorf("unmarshal int32: value %d out of range", dec)
	}

	*n = NullInt32{Valid: true, Int32: int32(dec)}

	return nil
}

// MarshalCQL - marshaller for cql
func (n NullInt64) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return marshalIntLike(info, n.Int64)
}

// UnmarshalCQL - unmarshaller for cql
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

// MarshalCQL - marshaller for cql
func (n NullString) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return []byte(n.String), nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullString) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if data == nil {
		*n = NullString{}
		return nil
	}

	*n = NullString{Valid: true, String: string(data)}

	return nil
}

// MarshalCQL - marshaller for cql
func (n NullTime) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	if n.Time.IsZero() {
		return []byte{}, nil
	}

	// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L1160
	x := int64(n.Time.UTC().Unix()*1e3) + int64(n.Time.UTC().Nanosecond()/1e6)

	enc := encBigInt(x)

	return enc, nil
}

// UnmarshalCQL - unmarshaller for cql
func (n *NullTime) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	switch {
	case data == nil:
		*n = NullTime{}
		return nil
	case len(data) == 0:
		*n = NullTime{Valid: true}
		return nil
	case len(data) != 8:
		*n = NullTime{}
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
