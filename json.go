package nan

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.String)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		*n = NullString{}
		return nil
	}

	var dec string

	err := jsoniter.Unmarshal(b, &dec)
	if err != nil {
		return err
	}

	*n = NullString{String: dec, Valid: true}

	return nil
}
