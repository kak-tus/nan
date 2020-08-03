package nan

import (
	"testing"
	"time"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

type cqlNanCoder interface {
	MarshalCQL(gocql.TypeInfo) ([]byte, error)
	UnmarshalCQL(gocql.TypeInfo, []byte) error
}

func doCQLTest(t *testing.T, cType gocql.Type, val1, val2 interface{}, nanVal1, nanVal2 cqlNanCoder) {
	cTypeEnc := gocql.NewNativeType(5, cType, "")

	encodedNan, err := nanVal1.MarshalCQL(cTypeEnc)
	assert.NoError(t, err)

	encodedVal, err := gocql.Marshal(cTypeEnc, val1)
	assert.NoError(t, err)

	assert.Equal(t, encodedNan, encodedVal)

	err = nanVal2.UnmarshalCQL(cTypeEnc, encodedNan)
	assert.NoError(t, err)

	assert.Equal(t, nanVal1, nanVal2)

	// not decode nil values
	if val2 != nil {
		err = gocql.Unmarshal(cTypeEnc, encodedVal, val2)
		assert.NoError(t, err)

		assert.Equal(t, val1, val2)
	}
}

func TestCQLNullBool(t *testing.T) {
	v1, v2 := true, true
	doCQLTest(t, gocql.TypeBoolean, &v1, &v2, &NullBool{Bool: v1, Valid: true}, &NullBool{})

	v1, v2 = false, false
	doCQLTest(t, gocql.TypeBoolean, &v1, &v2, &NullBool{Bool: false, Valid: true}, &NullBool{})

	doCQLTest(t, gocql.TypeBoolean, nil, nil, &NullBool{Valid: false}, &NullBool{})
	doCQLTest(t, gocql.TypeBoolean, nil, nil, &NullBool{Valid: false}, &NullBool{Valid: true})
}

func TestCQLNullFloat32(t *testing.T) {
	v1, v2 := float32(7676.76), float32(7676.76)
	doCQLTest(t, gocql.TypeFloat, &v1, &v2, &NullFloat32{Float32: v1, Valid: true}, &NullFloat32{})

	doCQLTest(t, gocql.TypeFloat, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{})
	doCQLTest(t, gocql.TypeFloat, nil, nil, &NullFloat32{Valid: false}, &NullFloat32{Valid: true})
}

func TestCQLNullFloat64(t *testing.T) {
	v1, v2 := float64(7676.76), float64(7676.76)
	doCQLTest(t, gocql.TypeDouble, &v1, &v2, &NullFloat64{Float64: v1, Valid: true}, &NullFloat64{})

	doCQLTest(t, gocql.TypeDouble, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{})
	doCQLTest(t, gocql.TypeDouble, nil, nil, &NullFloat64{Valid: false}, &NullFloat64{Valid: true})
}

func TestCQLNullInt(t *testing.T) {
	v1, v2 := int(7676), int(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	v1, v2 = int(76), int(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt{Int: v1, Valid: true}, &NullInt{})

	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt{Valid: false}, &NullInt{})
	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt{Valid: false}, &NullInt{Valid: true})
}

func TestCQLNullInt8(t *testing.T) {
	v1, v2 := int8(76), int8(76)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt8{Int8: v1, Valid: true}, &NullInt8{})

	doCQLTest(t, gocql.TypeTinyInt, nil, nil, &NullInt8{Valid: false}, &NullInt8{})
	doCQLTest(t, gocql.TypeTinyInt, nil, nil, &NullInt8{Valid: false}, &NullInt8{Valid: true})
}

func TestCQLNullInt16(t *testing.T) {
	v1, v2 := int16(7676), int16(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	v1, v2 = int16(76), int16(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt16{Int16: v1, Valid: true}, &NullInt16{})

	doCQLTest(t, gocql.TypeSmallInt, nil, nil, &NullInt16{Valid: false}, &NullInt16{})
	doCQLTest(t, gocql.TypeSmallInt, nil, nil, &NullInt16{Valid: false}, &NullInt16{Valid: true})
}

func TestCQLNullInt32(t *testing.T) {
	v1, v2 := int32(7676), int32(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	v1, v2 = int32(76), int32(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt32{Int32: v1, Valid: true}, &NullInt32{})

	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt32{Valid: false}, &NullInt32{})
	doCQLTest(t, gocql.TypeInt, nil, nil, &NullInt32{Valid: false}, &NullInt32{Valid: true})
}

func TestCQLNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doCQLTest(t, gocql.TypeBigInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})
	doCQLTest(t, gocql.TypeInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})
	doCQLTest(t, gocql.TypeSmallInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	v1, v2 = int64(76), int64(76)
	doCQLTest(t, gocql.TypeTinyInt, &v1, &v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doCQLTest(t, gocql.TypeBigInt, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doCQLTest(t, gocql.TypeBigInt, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})
}

func TestCQLNullString(t *testing.T) {
	v1, v2 := "7676", "7676"
	doCQLTest(t, gocql.TypeVarchar, &v1, &v2, &NullString{String: v1, Valid: true}, &NullString{})

	v1, v2 = "", ""
	doCQLTest(t, gocql.TypeVarchar, &v1, &v2, &NullString{String: v1, Valid: true}, &NullString{})

	doCQLTest(t, gocql.TypeVarchar, nil, nil, &NullString{Valid: false}, &NullString{})
	doCQLTest(t, gocql.TypeVarchar, nil, nil, &NullString{Valid: false}, &NullString{Valid: true})
}

func TestCQLNullTime(t *testing.T) {
	// Cassandra truncates it
	v1 := time.Now().UTC().Truncate(time.Millisecond)
	doCQLTest(t, gocql.TypeTimestamp, &v1, &time.Time{}, &NullTime{Time: v1, Valid: true}, &NullTime{})

	v1 = time.Time{}
	doCQLTest(t, gocql.TypeTimestamp, &v1, &time.Time{}, &NullTime{Valid: true}, &NullTime{})

	doCQLTest(t, gocql.TypeTimestamp, nil, nil, &NullTime{Valid: false}, &NullTime{})
	doCQLTest(t, gocql.TypeTimestamp, nil, nil, &NullTime{Valid: false}, &NullTime{Valid: true})
}
