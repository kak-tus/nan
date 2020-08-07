package nan

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeprecatedHelpers(t *testing.T) {
	assert.Equal(t, NullString{String: "1", Valid: true}, StringToNullString("1"))
	assert.Equal(t, NullBool{Bool: true, Valid: true}, BoolToNullBool(true))
	assert.Equal(t, NullInt64{Int64: 1, Valid: true}, Int64ToNullInt64(1))

	tm := time.Now()
	assert.Equal(t, NullTime{Time: tm, Valid: true}, TimeToNullTime(tm))

	assert.Equal(t, NullFloat64{Float64: 1, Valid: true}, Float64ToNullFloat64(1))
	assert.Equal(t, NullFloat32{Float32: 1, Valid: true}, Float32ToNullFloat32(1))
	assert.Equal(t, NullInt8{Int8: 1, Valid: true}, Int8ToNullInt8(1))
	assert.Equal(t, NullInt16{Int16: 1, Valid: true}, Int16ToNullInt16(1))
	assert.Equal(t, NullInt32{Int32: 1, Valid: true}, Int32ToNullInt32(1))
	assert.Equal(t, NullInt{Int: 1, Valid: true}, IntToNullInt(1))
}

func TestHelpers(t *testing.T) {
	assert.Equal(t, NullString{String: "1", Valid: true}, String("1"))
	assert.Equal(t, NullBool{Bool: true, Valid: true}, Bool(true))
	assert.Equal(t, NullInt64{Int64: 1, Valid: true}, Int64(1))

	tm := time.Now()
	assert.Equal(t, NullTime{Time: tm, Valid: true}, Time(tm))

	assert.Equal(t, NullFloat64{Float64: 1, Valid: true}, Float64(1))
	assert.Equal(t, NullFloat32{Float32: 1, Valid: true}, Float32(1))
	assert.Equal(t, NullInt8{Int8: 1, Valid: true}, Int8(1))
	assert.Equal(t, NullInt16{Int16: 1, Valid: true}, Int16(1))
	assert.Equal(t, NullInt32{Int32: 1, Valid: true}, Int32(1))
	assert.Equal(t, NullInt{Int: 1, Valid: true}, Int(1))
}
