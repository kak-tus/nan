package nan

import (
	"database/sql"
	"database/sql/driver"
	"testing"

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
