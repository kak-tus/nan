package nan

import (
	"testing"
	"time"

	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/assert"
)

func TestNanValidator(t *testing.T) {
	nullValues := []Validator{
		NullInt64{},
		NullInt32{},
		NullInt16{},
		NullInt8{},
		NullInt{},
		NullBool{},
		NullFloat64{},
		NullFloat32{},
		NullString{},
		NullTime{},
	}
	nonnullvalues := []Validator{
		NullInt64{Valid: true, Int64: 1},
		NullInt32{Valid: true, Int32: 1},
		NullInt16{Valid: true, Int16: 1},
		NullInt8{Valid: true, Int8: 1},
		NullInt{Valid: true, Int: 1},
		NullBool{Valid: true, Bool: true},
		NullFloat64{Valid: true, Float64: 1},
		NullFloat32{Valid: true, Float32: 1},
		NullString{Valid: true, String: "1"},
		NullTime{Valid: true, Time: time.Now()},
	}
	for _, v := range nullValues {
		assert.False(t, v.IsValid())
	}
	for _, v := range nonnullvalues {
		assert.True(t, v.IsValid())
	}
}

func TestNanDefined(t *testing.T) {
	nullValues := []easyjson.Optional{
		NullInt64{},
		NullInt32{},
		NullInt16{},
		NullInt8{},
		NullInt{},
		NullBool{},
		NullFloat64{},
		NullFloat32{},
		NullString{},
		NullTime{},
	}
	nonnullvalues := []easyjson.Optional{
		NullInt64{Valid: true, Int64: 1},
		NullInt32{Valid: true, Int32: 1},
		NullInt16{Valid: true, Int16: 1},
		NullInt8{Valid: true, Int8: 1},
		NullInt{Valid: true, Int: 1},
		NullBool{Valid: true, Bool: true},
		NullFloat64{Valid: true, Float64: 1},
		NullFloat32{Valid: true, Float32: 1},
		NullString{Valid: true, String: "1"},
		NullTime{Valid: true, Time: time.Now()},
	}
	for _, v := range nullValues {
		assert.False(t, v.IsDefined())
	}
	for _, v := range nonnullvalues {
		assert.True(t, v.IsDefined())
	}
}
