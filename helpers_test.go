package nan

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDeprecatedHelpers(t *testing.T) {
	require.Equal(t, NullString{String: "1", Valid: true}, StringToNullString("1"))
	require.Equal(t, NullBool{Bool: true, Valid: true}, BoolToNullBool(true))
	require.Equal(t, NullInt64{Int64: 1, Valid: true}, Int64ToNullInt64(1))

	tm := time.Now()
	require.Equal(t, NullTime{Time: tm, Valid: true}, TimeToNullTime(tm))

	require.Equal(t, NullFloat64{Float64: 1, Valid: true}, Float64ToNullFloat64(1))
	require.Equal(t, NullFloat32{Float32: 1, Valid: true}, Float32ToNullFloat32(1))
	require.Equal(t, NullInt8{Int8: 1, Valid: true}, Int8ToNullInt8(1))
	require.Equal(t, NullInt16{Int16: 1, Valid: true}, Int16ToNullInt16(1))
	require.Equal(t, NullInt32{Int32: 1, Valid: true}, Int32ToNullInt32(1))
	require.Equal(t, NullInt{Int: 1, Valid: true}, IntToNullInt(1))
}

func TestHelpers(t *testing.T) {
	require.Equal(t, NullString{String: "1", Valid: true}, String("1"))
	require.Equal(t, NullBool{Bool: true, Valid: true}, Bool(true))
	require.Equal(t, NullInt64{Int64: 1, Valid: true}, Int64(1))

	tm := time.Now()
	require.Equal(t, NullTime{Time: tm, Valid: true}, Time(tm))

	require.Equal(t, NullFloat64{Float64: 1, Valid: true}, Float64(1))
	require.Equal(t, NullFloat32{Float32: 1, Valid: true}, Float32(1))
	require.Equal(t, NullInt8{Int8: 1, Valid: true}, Int8(1))
	require.Equal(t, NullInt16{Int16: 1, Valid: true}, Int16(1))
	require.Equal(t, NullInt32{Int32: 1, Valid: true}, Int32(1))
	require.Equal(t, NullInt{Int: 1, Valid: true}, Int(1))

	require.Equal(t, NullUint{Uint: 1, Valid: true}, Uint(1))
	require.Equal(t, NullUint8{Uint8: 1, Valid: true}, Uint8(1))
	require.Equal(t, NullUint16{Uint16: 1, Valid: true}, Uint16(1))
	require.Equal(t, NullUint32{Uint32: 1, Valid: true}, Uint32(1))
	require.Equal(t, NullUint64{Uint64: 1, Valid: true}, Uint64(1))
}
