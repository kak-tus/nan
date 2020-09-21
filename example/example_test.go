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

	encJsoniter, err := jsoniter.Marshal(val)
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
}
