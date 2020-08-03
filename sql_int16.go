package nan

import "database/sql/driver"

// Scan - scan value from sql driver
func (n *NullInt16) Scan(value interface{}) error {
	if value == nil {
		n.Int16, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int16, value)
}

// Value implements the driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int16), nil
}
