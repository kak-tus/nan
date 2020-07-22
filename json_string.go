package nan

import (
	"bytes"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

func init() {
	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullString)(ptr)) = NullString{String: iter.ReadString(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullString)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteString(t.String)
		},
		nil,
	)
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.String)
}

func (n *NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullString{}
		return nil
	}

	var res string

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullString{String: res, Valid: true}

	return nil
}

func (n NullString) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.String(n.String)
}

func (n *NullString) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullString{}

		in.Skip()

		return
	}

	*n = NullString{String: in.String(), Valid: true}
}
