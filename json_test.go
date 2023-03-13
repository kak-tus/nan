package nan

import (
	json "encoding/json"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type jsonNanCoder interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	MarshalEasyJSON(out *jwriter.Writer)
	UnmarshalEasyJSON(in *jlexer.Lexer)
}

func doJSONTest(t *testing.T, val1, val2 interface{}, nanVal1, nanVal2 jsonNanCoder) {
	encodedNanJSON, err := json.Marshal(nanVal1)
	assert.NoError(t, err)

	encodedValJSON, err := json.Marshal(val1)
	assert.NoError(t, err)

	assert.Equal(t, encodedNanJSON, encodedValJSON)

	encodedValJsoniter, err := jsoniter.Marshal(nanVal1)
	assert.NoError(t, err)

	assert.Equal(t, encodedNanJSON, encodedValJsoniter)

	w := jwriter.Writer{}

	nanVal1.MarshalEasyJSON(&w)
	assert.NoError(t, w.Error)

	assert.Equal(t, encodedNanJSON, w.Buffer.BuildBytes())

	err = json.Unmarshal(encodedNanJSON, nanVal2)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)

	err = jsoniter.Unmarshal(encodedNanJSON, nanVal2)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)

	r := jlexer.Lexer{Data: encodedNanJSON}

	nanVal2.UnmarshalEasyJSON(&r)
	assert.NoError(t, r.Error())

	assert.Equal(t, nanVal1, nanVal2)
}

func TestJSONNullBool(t *testing.T) {
	v1, v2 := true, true
	doJSONTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doJSONTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	doJSONTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doJSONTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})

	v := NullBool{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doJSONTest(t, v1, v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})

	v := NullFloat32{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doJSONTest(t, v1, v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})

	v := NullFloat64{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doJSONTest(t, v1, v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})

	v := NullInt{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doJSONTest(t, v1, v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})

	v := NullInt8{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doJSONTest(t, v1, v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doJSONTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doJSONTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})

	v := NullInt16{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doJSONTest(t, v1, v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})

	v := NullInt32{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doJSONTest(t, v1, v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doJSONTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doJSONTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})

	v := NullInt64{}
	assert.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	doJSONTest(t, nil, nil, &NullString{Valid: false}, &NullString{})

	v := NullString{}
	assert.Error(t, json.Unmarshal([]byte("123"), &v))
}

func TestJSONNullTime(t *testing.T) {
	v1 := time.Now().UTC()
	doJSONTest(t, v1, &time.Time{}, &NullTime{Time: v1, Valid: true}, &NullTime{})

	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})

	err := jsoniter.Unmarshal([]byte(""), &NullTime{})
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte("\"wrong\""), &NullTime{})
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte("\"wrong\""), &NullTime{})
	assert.Error(t, err)

	err = json.Unmarshal([]byte("\"wrong\""), &NullTime{})
	assert.Error(t, err)

	err = json.Unmarshal([]byte("123"), &NullTime{})
	assert.Error(t, err)

	r := jlexer.Lexer{Data: []byte("\"wrong\"")}

	var nanVal NullTime

	nanVal.UnmarshalEasyJSON(&r)
	assert.Error(t, r.Error())
}

func TestOmitempty(t *testing.T) {
	var decoded struct {
		Bool    NullBool    `json:"bool,omitempty"`
		Float32 NullFloat32 `json:"float32,omitempty"`
		Float64 NullFloat64 `json:"float64,omitempty"`
		Int     NullInt     `json:"int,omitempty"`
		Int8    NullInt8    `json:"int8,omitempty"`
		Int16   NullInt16   `json:"int16,omitempty"`
		Int32   NullInt32   `json:"int32,omitempty"`
		Int64   NullInt64   `json:"int64,omitempty"`
		String  NullString  `json:"string,omitempty"`
		Time    NullTime    `json:"time,omitempty"`
	}

	encoded, err := jsoniter.MarshalToString(decoded)
	assert.NoError(t, err)

	assert.Equal(t, "{}", encoded)
}

func TestJSONNullUint(t *testing.T) {
	v1, v2 := uint(7676), uint(7676)
	doJSONTest(t, v1, v2, &NullUint{Uint: v1, Valid: true}, &NullUint{})

	doJSONTest(t, nil, nil, &NullUint{Valid: false}, &NullUint{})
	doJSONTest(t, nil, nil, &NullUint{Valid: false}, &NullUint{Valid: true})

	v := NullUint{}
	require.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullUint8(t *testing.T) {
	v1, v2 := uint8(76), uint8(76)
	doJSONTest(t, v1, v2, &NullUint8{Uint8: v1, Valid: true}, &NullUint8{})

	doJSONTest(t, nil, nil, &NullUint8{Valid: false}, &NullUint8{})
	doJSONTest(t, nil, nil, &NullUint8{Valid: false}, &NullUint8{Valid: true})

	v := NullUint8{}
	require.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullUint16(t *testing.T) {
	v1, v2 := uint16(7676), uint16(7676)
	doJSONTest(t, v1, v2, &NullUint16{Uint16: v1, Valid: true}, &NullUint16{})

	doJSONTest(t, nil, nil, &NullUint16{Valid: false}, &NullUint16{})
	doJSONTest(t, nil, nil, &NullUint16{Valid: false}, &NullUint16{Valid: true})

	v := NullUint16{}
	require.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullUint32(t *testing.T) {
	v1, v2 := uint32(7676), uint32(7676)
	doJSONTest(t, v1, v2, &NullUint32{Uint32: v1, Valid: true}, &NullUint32{})

	doJSONTest(t, nil, nil, &NullUint32{Valid: false}, &NullUint32{})
	doJSONTest(t, nil, nil, &NullUint32{Valid: false}, &NullUint32{Valid: true})

	v := NullUint32{}
	require.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}

func TestJSONNullUint64(t *testing.T) {
	v1, v2 := uint64(7676), uint64(7676)
	doJSONTest(t, v1, v2, &NullUint64{Uint64: v1, Valid: true}, &NullUint64{})

	doJSONTest(t, nil, nil, &NullUint64{Valid: false}, &NullUint64{})
	doJSONTest(t, nil, nil, &NullUint64{Valid: false}, &NullUint64{Valid: true})

	v := NullUint64{}
	require.Error(t, json.Unmarshal([]byte("\"wrong\""), &v))
}
