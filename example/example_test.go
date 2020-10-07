package example

import (
	json "encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestStructs(t *testing.T) {
	val := MyStruct2{
		Val3: NullMyStruct{MyStruct: MyStruct{ID: 1}, Valid: true},
		Val4: MyStruct{ID: 1},
	}

	encJSON, err := json.Marshal(val)
	assert.NoError(t, err)
	fmt.Println(string(encJSON))

	// use separate struct without easyjson coders
	// Because easyjson generates marshallers for encode/decode and original
	// jsoniter registered functions not used
	valJsoniter := MyStruct3{
		Val3: NullMyStruct{MyStruct: MyStruct{ID: 1}, Valid: true},
		Val4: MyStruct{ID: 1},
	}

	encJsoniter, err := jsoniter.Marshal(valJsoniter)
	assert.NoError(t, err)
	fmt.Println(string(encJsoniter))

	assert.Equal(t, encJSON, encJsoniter)

	encEasyJSON, err := val.MarshalJSON()
	assert.NoError(t, err)
	fmt.Println(string(encEasyJSON))

	assert.Equal(t, encJSON, encEasyJSON)

	var targetJSON MyStruct2
	err = json.Unmarshal(encJSON, &targetJSON)
	assert.NoError(t, err)
	assert.Equal(t, val, targetJSON)

	var targetJsoniter MyStruct3
	err = jsoniter.Unmarshal(encJsoniter, &targetJsoniter)
	assert.NoError(t, err)
	assert.Equal(t, valJsoniter, targetJsoniter)

	var targetEasyjson MyStruct2
	err = targetEasyjson.UnmarshalJSON(encEasyJSON)
	assert.NoError(t, err)
	assert.Equal(t, val, targetEasyjson)
}
