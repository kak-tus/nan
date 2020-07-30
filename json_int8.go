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
		"nan.NullInt8",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt8)(ptr)) = NullInt8{Int8: iter.ReadInt8(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt8",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt8)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt8(t.Int8)
		},
		nil,
	)
}

func (n NullInt8) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Int8)
}

func (n *NullInt8) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt8{}
		return nil
	}

	var res int8

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt8{Int8: res, Valid: true}

	return nil
}

func (n NullInt8) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int8(n.Int8)
}

func (n *NullInt8) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt8{}

		in.Skip()

		return
	}

	*n = NullInt8{Int8: in.Int8(), Valid: true}
}
