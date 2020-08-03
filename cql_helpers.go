package nan

import (
	"fmt"
	"math"

	"github.com/gocql/gocql"
)

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L519
func encShort(x int16) []byte {
	p := make([]byte, 2)
	p[0] = byte(x >> 8)
	p[1] = byte(x)

	return p
}

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L526
func decShort(p []byte) int16 {
	if len(p) != 2 {
		return 0
	}

	return int16(p[0])<<8 | int16(p[1])
}

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L508
func encInt(x int32) []byte {
	return []byte{byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)}
}

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L512
func decInt(x []byte) int32 {
	if len(x) != 4 {
		return 0
	}

	return int32(x[0])<<24 | int32(x[1])<<16 | int32(x[2])<<8 | int32(x[3])
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

// https://github.com/gocql/gocql/blob/b454769479c6201d26d2a2d7a29ac9a0e6fbc9fc/marshal.go#L533
func decTiny(p []byte) int8 {
	if len(p) != 1 {
		return 0
	}

	return int8(p[0])
}

func cqlMarshalErrorf(format string, args ...interface{}) gocql.MarshalError {
	return gocql.MarshalError(fmt.Sprintf(format, args...))
}

func marshalIntLike(info gocql.TypeInfo, value int64) ([]byte, error) {
	switch info.Type() {
	case gocql.TypeTinyInt:
		if value > math.MaxInt8 || value < math.MinInt8 {
			return nil, cqlMarshalErrorf("marshal tinyint: value %d out of range", value)
		}
		return []byte{byte(value)}, nil
	case gocql.TypeSmallInt:
		if value > math.MaxInt16 || value < math.MinInt16 {
			return nil, cqlMarshalErrorf("marshal smallint: value %d out of range", value)
		}
		return encShort(int16(value)), nil
	case gocql.TypeInt:
		if value > math.MaxInt32 || value < math.MinInt32 {
			return nil, cqlMarshalErrorf("marshal int: value %d out of range", value)
		}
		return encInt(int32(value)), nil
	case gocql.TypeBigInt:
		return encBigInt(value), nil
	default:
		return nil, cqlMarshalErrorf("marshal %v: cannot marshal from integer", info.Type())
	}
}

func unmarshalIntLike(info gocql.TypeInfo, data []byte) (int64, error) {
	switch info.Type() {
	case gocql.TypeTinyInt:
		return int64(decTiny(data)), nil
	case gocql.TypeSmallInt:
		return int64(decShort(data)), nil
	case gocql.TypeInt:
		return int64(decInt(data)), nil
	case gocql.TypeBigInt:
		return decBigInt(data), nil
	default:
		return 0, cqlMarshalErrorf("unmarshal %v: cannot unmarshal into integer", info.Type())
	}
}
