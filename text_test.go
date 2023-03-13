package nan

import (
	"encoding"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type textNanEncoder interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	Validator
}

func doTextTest(t *testing.T, val interface{}, nanVal textNanEncoder, nanEmpty textNanEncoder, strVal string) {
	str, err := nanVal.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, []byte(strVal), str)

	err = nanEmpty.UnmarshalText(str)
	assert.NoError(t, err)
	if val == nil {
		assert.False(t, nanEmpty.IsValid())
	} else {
		assert.True(t, nanEmpty.IsValid())
		assert.Equal(t, nanVal, nanEmpty)
	}
}

func TestTextNullBool(t *testing.T) {
	doTextTest(t, true, &NullBool{Bool: true, Valid: true}, &NullBool{}, "true")
	doTextTest(t, false, &NullBool{Bool: false, Valid: true}, &NullBool{}, "false")
	doTextTest(t, nil, &NullBool{Valid: false}, &NullBool{}, "null")

	v := NullBool{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullFloat32(t *testing.T) {
	doTextTest(t, true, &NullFloat32{Float32: 1.5, Valid: true}, &NullFloat32{}, "1.5")
	doTextTest(t, nil, &NullFloat32{Valid: false}, &NullFloat32{}, "null")

	v := NullFloat32{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullFloat64(t *testing.T) {
	doTextTest(t, true, &NullFloat64{Float64: 1.5, Valid: true}, &NullFloat64{}, "1.5")
	doTextTest(t, nil, &NullFloat64{Valid: false}, &NullFloat64{}, "null")

	v := NullFloat64{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullInt(t *testing.T) {
	doTextTest(t, true, &NullInt{Int: math.MaxInt32, Valid: true}, &NullInt{}, strconv.Itoa(math.MaxInt32))
	doTextTest(t, nil, &NullInt{Valid: false}, &NullInt{}, "null")

	v := NullInt{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullInt8(t *testing.T) {
	doTextTest(t, true, &NullInt8{Int8: math.MaxInt8, Valid: true}, &NullInt8{}, strconv.Itoa(math.MaxInt8))
	doTextTest(t, nil, &NullInt8{Valid: false}, &NullInt8{}, "null")

	v := NullInt8{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullInt16(t *testing.T) {
	doTextTest(t, true, &NullInt16{Int16: math.MaxInt16, Valid: true}, &NullInt16{}, strconv.Itoa(math.MaxInt16))
	doTextTest(t, nil, &NullInt16{Valid: false}, &NullInt16{}, "null")

	v := NullInt16{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullInt32(t *testing.T) {
	doTextTest(t, true, &NullInt32{Int32: math.MaxInt32, Valid: true}, &NullInt32{}, strconv.Itoa(math.MaxInt32))
	doTextTest(t, nil, &NullInt32{Valid: false}, &NullInt32{}, "null")

	v := NullInt32{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullInt64(t *testing.T) {
	doTextTest(t, true, &NullInt64{Int64: math.MaxInt64, Valid: true}, &NullInt64{}, strconv.FormatInt(math.MaxInt64, 10))
	doTextTest(t, nil, &NullInt64{Valid: false}, &NullInt64{}, "null")

	v := NullInt64{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullString(t *testing.T) {
	doTextTest(t, true, &NullString{String: "str", Valid: true}, &NullString{}, `"str"`)
	doTextTest(t, true, &NullString{String: "null", Valid: true}, &NullString{}, `"null"`)
	doTextTest(t, nil, &NullString{Valid: false}, &NullString{}, "null")

	v := NullString{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullTime(t *testing.T) {
	tt := time.Now().Truncate(time.Second).UTC()
	doTextTest(t, true, &NullTime{Time: tt, Valid: true}, &NullTime{}, `"`+tt.Format(time.RFC3339Nano)+`"`)
	doTextTest(t, nil, &NullTime{Valid: false}, &NullTime{}, "null")

	v := NullTime{}
	assert.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullUint(t *testing.T) {
	max := strconv.FormatUint(math.MaxUint, 10)
	doTextTest(t, true, &NullUint{Uint: math.MaxUint, Valid: true}, &NullUint{}, max)
	doTextTest(t, nil, &NullUint{Valid: false}, &NullUint{}, "null")

	v := NullUint{}
	require.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullUint8(t *testing.T) {
	max := strconv.FormatUint(math.MaxUint8, 10)
	doTextTest(t, true, &NullUint8{Uint8: math.MaxUint8, Valid: true}, &NullUint8{}, max)
	doTextTest(t, nil, &NullUint8{Valid: false}, &NullUint8{}, "null")

	v := NullUint8{}
	require.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullUint16(t *testing.T) {
	max := strconv.FormatUint(math.MaxUint16, 10)
	doTextTest(t, true, &NullUint16{Uint16: math.MaxUint16, Valid: true}, &NullUint16{}, max)
	doTextTest(t, nil, &NullUint16{Valid: false}, &NullUint16{}, "null")

	v := NullUint16{}
	require.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullUint32(t *testing.T) {
	max := strconv.FormatUint(math.MaxUint32, 10)
	doTextTest(t, true, &NullUint32{Uint32: math.MaxUint32, Valid: true}, &NullUint32{}, max)
	doTextTest(t, nil, &NullUint32{Valid: false}, &NullUint32{}, "null")

	v := NullUint32{}
	require.Error(t, v.UnmarshalText([]byte("wrong")))
}

func TestTextNullUint64(t *testing.T) {
	max := strconv.FormatUint(math.MaxUint64, 10)
	doTextTest(t, true, &NullUint64{Uint64: math.MaxUint64, Valid: true}, &NullUint64{}, max)
	doTextTest(t, nil, &NullUint64{Valid: false}, &NullUint64{}, "null")

	v := NullUint64{}
	require.Error(t, v.UnmarshalText([]byte("wrong")))
}
