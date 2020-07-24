package nan

import (
	json "encoding/json"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	"github.com/stretchr/testify/assert"
)

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
