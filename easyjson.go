package nan

import (
	"time"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// MarshalEasyJSON - marshaller for easyjson
func (n NullBool) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Bool(n.Bool)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullBool) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullBool{}

		in.Skip()

		return
	}

	*n = NullBool{Bool: in.Bool(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullFloat32) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Float32(n.Float32)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullFloat32) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullFloat32{}

		in.Skip()

		return
	}

	*n = NullFloat32{Float32: in.Float32(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullFloat64) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Float64(n.Float64)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullFloat64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullFloat64{}

		in.Skip()

		return
	}

	*n = NullFloat64{Float64: in.Float64(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullInt) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int(n.Int)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullInt) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt{}

		in.Skip()

		return
	}

	*n = NullInt{Int: in.Int(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullInt8) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int8(n.Int8)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullInt8) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt8{}

		in.Skip()

		return
	}

	*n = NullInt8{Int8: in.Int8(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullInt16) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int16(n.Int16)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullInt16) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt16{}

		in.Skip()

		return
	}

	*n = NullInt16{Int16: in.Int16(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullInt32) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int32(n.Int32)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullInt32) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt32{}

		in.Skip()

		return
	}

	*n = NullInt32{Int32: in.Int32(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullInt64) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int64(n.Int64)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullInt64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt64{}

		in.Skip()

		return
	}

	*n = NullInt64{Int64: in.Int64(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullString) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.String(n.String)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullString) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullString{}

		in.Skip()

		return
	}

	*n = NullString{String: in.String(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullTime) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.String(n.Time.Format(time.RFC3339Nano))
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullTime) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullTime{}

		in.Skip()

		return
	}

	res, err := time.Parse(time.RFC3339Nano, in.String())
	if err != nil {
		in.AddError(err)
		return
	}

	*n = NullTime{Time: res, Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullUint) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Uint(n.Uint)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullUint) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullUint{}

		in.Skip()

		return
	}

	*n = NullUint{Uint: in.Uint(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullUint8) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Uint8(n.Uint8)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullUint8) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullUint8{}

		in.Skip()

		return
	}

	*n = NullUint8{Uint8: in.Uint8(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullUint16) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Uint16(n.Uint16)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullUint16) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullUint16{}

		in.Skip()

		return
	}

	*n = NullUint16{Uint16: in.Uint16(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullUint32) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Uint32(n.Uint32)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullUint32) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullUint32{}

		in.Skip()

		return
	}

	*n = NullUint32{Uint32: in.Uint32(), Valid: true}
}

// MarshalEasyJSON - marshaller for easyjson
func (n NullUint64) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Uint64(n.Uint64)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *NullUint64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullUint64{}

		in.Skip()

		return
	}

	*n = NullUint64{Uint64: in.Uint64(), Valid: true}
}

func (n initialTemplateType) MarshalEasyJSON(out *jwriter.Writer) {
	// Function only for code validity
}

func (n *initialTemplateType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	// Function only for code validity
}

func (n initialTemplateType) MarshalJSON() ([]byte, error) {
	// Function only for code validity
	return nil, nil
}

func (n *initialTemplateType) UnmarshalJSON(data []byte) error {
	// Function only for code validity
	return nil
}
