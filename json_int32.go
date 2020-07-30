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
		"nan.NullInt32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt32)(ptr)) = NullInt32{Int32: iter.ReadInt32(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt32(t.Int32)
		},
		nil,
	)
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Int32)
}

func (n *NullInt32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt32{}
		return nil
	}

	var res int32

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt32{Int32: res, Valid: true}

	return nil
}

func (n NullInt32) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int32(n.Int32)
}

func (n *NullInt32) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt32{}

		in.Skip()

		return
	}

	*n = NullInt32{Int32: in.Int32(), Valid: true}
}
