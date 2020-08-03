package nan

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type sqlNanCoder interface {
	sql.Scanner
	Value() (driver.Value, error)
}

func doSQLTest(t *testing.T, val1, val2 interface{}, nanVal1, nanVal2 sqlNanCoder) {
	val, err := nanVal1.Value()
	assert.NoError(t, err)
	assert.EqualValues(t, val1, val) //sql.Null* types return int64 for all integer types

	assert.NoError(t, nanVal2.Scan(val2))
	val, err = nanVal2.Value()
	assert.NoError(t, err)
	assert.EqualValues(t, val2, val)
}

func TestSQLNullBool(t *testing.T) {
	v1, v2 := true, true
	doSQLTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doSQLTest(t, v1, v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	doSQLTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doSQLTest(t, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})
}

func TestSQLNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doSQLTest(t, v1, v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doSQLTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doSQLTest(t, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}

func TestSQLNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doSQLTest(t, v1, v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doSQLTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doSQLTest(t, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}

func TestSQLNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doSQLTest(t, v1, v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doSQLTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doSQLTest(t, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}

func TestSQLNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doSQLTest(t, v1, v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doSQLTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doSQLTest(t, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}

func TestSQLNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doSQLTest(t, v1, v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doSQLTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doSQLTest(t, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})
}

func TestSQLNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doSQLTest(t, v1, v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doSQLTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doSQLTest(t, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}

func TestSQLNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doSQLTest(t, v1, v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doSQLTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doSQLTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})
}

func TestSQLNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doSQLTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doSQLTest(t, v1, v2, &NullString{String: v1, Valid: true}, &NullString{})

	doSQLTest(t, nil, nil, &NullString{Valid: false}, &NullString{})
	doSQLTest(t, nil, nil, &NullString{Valid: false}, &NullString{Valid: true})
}

func TestSQLNullTime(t *testing.T) {
	v1, v2 := time.Now().UTC(), time.Now().UTC()
	doSQLTest(t, v1, v2, &NullTime{Time: v1, Valid: true}, &NullTime{})

	v1, v2 = time.Time{}, time.Time{}
	doSQLTest(t, v1, v2, &NullTime{Valid: true}, &NullTime{})

	doSQLTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doSQLTest(t, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})
}
