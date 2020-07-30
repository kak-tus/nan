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
		"nan.NullInt16",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt16)(ptr)) = NullInt16{Int16: iter.ReadInt16(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt16",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt16)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt16(t.Int16)
		},
		nil,
	)
}

func (n NullInt16) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Int16)
}

func (n *NullInt16) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt16{}
		return nil
	}

	var res int16

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt16{Int16: res, Valid: true}

	return nil
}

func (n NullInt16) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int16(n.Int16)
}

func (n *NullInt16) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt16{}

		in.Skip()

		return
	}

	*n = NullInt16{Int16: in.Int16(), Valid: true}
}
