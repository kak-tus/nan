package nan

import (
	"bytes"
	"encoding/json"
	"time"
)

// MarshalJSON - marshaller for json
func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Bool)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullBool{}
		return nil
	}

	var res bool

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullBool{Bool: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullFloat32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float32)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullFloat32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat32{}
		return nil
	}

	var res float32

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullFloat32{Float32: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float64)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat64{}
		return nil
	}

	var res float64

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullFloat64{Float64: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullInt) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt{}
		return nil
	}

	var res int

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt{Int: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullInt8) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int8)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt8) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt8{}
		return nil
	}

	var res int8

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt8{Int8: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullInt16) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int16)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt16{}
		return nil
	}

	var res int16

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt16{Int16: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int32)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt32{}
		return nil
	}

	var res int32

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt32{Int32: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int64)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt64{}
		return nil
	}

	var res int64

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt64{Int64: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.String)
}

// UnmarshalJSON - unmarshaller for json
func (n *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullString{}
		return nil
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullString{String: res, Valid: true}

	return nil
}

// MarshalJSON - marshaller for json
func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Time.Format(time.RFC3339Nano))
}

// UnmarshalJSON - unmarshaller for json
func (n *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullTime{}
		return nil
	}

	var dec string

	err := json.Unmarshal(data, &dec)
	if err != nil {
		return err
	}

	res, err := time.Parse(time.RFC3339Nano, string(dec))
	if err != nil {
		return err
	}

	*n = NullTime{Time: res, Valid: true}

	return nil
}
