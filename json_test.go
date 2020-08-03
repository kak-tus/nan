package nan

import (
	json "encoding/json"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/stretchr/testify/assert"
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
}

func TestJSONNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doJSONTest(t, v1, v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doJSONTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}

func TestJSONNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doJSONTest(t, v1, v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doJSONTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}

func TestJSONNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doJSONTest(t, v1, v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doJSONTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}

func TestJSONNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doJSONTest(t, v1, v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doJSONTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}

func TestJSONNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doJSONTest(t, v1, v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doJSONTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doJSONTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})
}

func TestJSONNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doJSONTest(t, v1, v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doJSONTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}

func TestJSONNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doJSONTest(t, v1, v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doJSONTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doJSONTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})
}

func TestJSONNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doJSONTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	doJSONTest(t, nil, nil, &NullString{Valid: false}, &NullString{})
}

func TestJSONNullTime(t *testing.T) {
	v1 := time.Now().UTC()
	doJSONTest(t, v1, &time.Time{}, &NullTime{Time: v1, Valid: true}, &NullTime{})

	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doJSONTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})

	err := jsoniter.Unmarshal([]byte(""), &NullTime{})
	assert.Error(t, err)

	err = jsoniter.Unmarshal([]byte("wrong"), &NullTime{})
	assert.Error(t, err)

	err = json.Unmarshal([]byte("wrong"), &NullTime{})
	assert.Error(t, err)

	r := jlexer.Lexer{Data: []byte("wrong")}

	var nanVal NullTime

	nanVal.UnmarshalEasyJSON(&r)
	assert.Error(t, r.Error())
}
