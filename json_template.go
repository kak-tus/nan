package nan

import (
	"bytes"
	"encoding/json"
)

// MarshalJSON - marshaller for json
func (n nullTemplateType) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.NullTemplateValue)
}

// UnmarshalJSON - unmarshaller for json
func (n *nullTemplateType) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = nullTemplateType{}
		return nil
	}

	var res initialTemplateType

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = nullTemplateType{NullTemplateValue: res, Valid: true}

	return nil
}
