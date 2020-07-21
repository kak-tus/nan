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

	err = json.Unmarshal(encodedNanJSON, nanVal2)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)
}
