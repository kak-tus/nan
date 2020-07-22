package nan

import (
	json "encoding/json"
	"testing"

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
