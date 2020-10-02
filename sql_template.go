package nan

import "database/sql/driver"

// Value implements the driver Valuer interface.
func (n nullTemplateType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.NullTemplateValue.Value()
}

// Scan - scan value from sql driver
func (n *nullTemplateType) Scan(value interface{}) error {
	if value == nil {
		*n = nullTemplateType{}
		return nil
	}

	var val initialTemplateType
	if err := val.Scan(value); err != nil {
		return err
	}

	*n = nullTemplateType{Valid: true, NullTemplateValue: val}

	return nil
}
