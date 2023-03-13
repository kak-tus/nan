package nan

import (
	"testing"
	"time"

	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/require"
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
		NullUint64{},
		NullUint32{},
		NullUint16{},
		NullUint8{},
		NullUint{},
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
		NullUint64{Valid: true, Uint64: 1},
		NullUint32{Valid: true, Uint32: 1},
		NullUint16{Valid: true, Uint16: 1},
		NullUint8{Valid: true, Uint8: 1},
		NullUint{Valid: true, Uint: 1},
	}

	for _, v := range nullValues {
		require.False(t, v.IsValid())
	}

	for _, v := range nonnullvalues {
		require.True(t, v.IsValid())
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
		NullUint64{},
		NullUint32{},
		NullUint16{},
		NullUint8{},
		NullUint{},
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
		NullUint64{Valid: true, Uint64: 1},
		NullUint32{Valid: true, Uint32: 1},
		NullUint16{Valid: true, Uint16: 1},
		NullUint8{Valid: true, Uint8: 1},
		NullUint{Valid: true, Uint: 1},
	}

	for _, v := range nullValues {
		require.False(t, v.IsDefined())
	}

	for _, v := range nonnullvalues {
		require.True(t, v.IsDefined())
	}
}
