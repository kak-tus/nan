package example

import (
	"database/sql/driver"

	"github.com/gocql/gocql"
)

type MyStruct struct {
	ID   int
	Name string
}

func (n MyStruct) Value() (driver.Value, error) {
	// This is only an example
	// Form normal data for sql
	return n.ID, nil
}

func (n *MyStruct) Scan(value interface{}) error {
	// This is only an example
	// Get normal data from value
	*n = MyStruct{ID: 1}

	return nil
}

func (n MyStruct) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	// This is only an example
	// Form normal data for cql
	return []byte(n.Name), nil
}

func (n *MyStruct) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	// This is only an example
	// Get normal data from value
	*n = MyStruct{ID: 1}

	return nil
}
