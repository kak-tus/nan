package nan

import (
	"bytes"
	"encoding/json"
	"time"
)

// MarshalText - marshaller for text
func (n NullBool) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Bool)
}

// UnmarshalText - unmarshaller for text
func (n *NullBool) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullBool{}
		return nil
	}

	var res bool

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullBool{Bool: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullFloat32) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Float32)
}

// UnmarshalText - unmarshaller for text
func (n *NullFloat32) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullFloat32{}
		return nil
	}

	var res float32

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullFloat32{Float32: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullFloat64) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Float64)
}

// UnmarshalText - unmarshaller for text
func (n *NullFloat64) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullFloat64{}
		return nil
	}

	var res float64

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullFloat64{Float64: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullInt) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int)
}

// UnmarshalText - unmarshaller for text
func (n *NullInt) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullInt{}
		return nil
	}

	var res int

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullInt{Int: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullInt8) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int8)
}

// UnmarshalText - unmarshaller for text
func (n *NullInt8) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullInt8{}
		return nil
	}

	var res int8

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullInt8{Int8: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullInt16) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int16)
}

// UnmarshalText - unmarshaller for text
func (n *NullInt16) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullInt16{}
		return nil
	}

	var res int16

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullInt16{Int16: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullInt32) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int32)
}

// UnmarshalText - unmarshaller for text
func (n *NullInt32) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullInt32{}
		return nil
	}

	var res int32

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullInt32{Int32: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullInt64) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int64)
}

// UnmarshalText - unmarshaller for text
func (n *NullInt64) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullInt64{}
		return nil
	}

	var res int64

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullInt64{Int64: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullString) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.String)
}

// UnmarshalText - unmarshaller for text
func (n *NullString) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullString{}
		return nil
	}

	var res string

	err := json.Unmarshal(text, &res)
	if err != nil {
		return err
	}

	*n = NullString{String: res, Valid: true}

	return nil
}

// MarshalText - marshaller for text
func (n NullTime) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Time.Format(time.RFC3339Nano))
}

// UnmarshalText - unmarshaller for text
func (n *NullTime) UnmarshalText(text []byte) error {
	if bytes.Equal(text, []byte("null")) {
		*n = NullTime{}
		return nil
	}

	var dec string

	err := json.Unmarshal(text, &dec)
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
