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
		"nan.NullInt",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt)(ptr)) = NullInt{Int: iter.ReadInt(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt(t.Int)
		},
		nil,
	)
}

func (n NullInt) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Int)
}

func (n *NullInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt{}
		return nil
	}

	var res int

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt{Int: res, Valid: true}

	return nil
}

func (n NullInt) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int(n.Int)
}

func (n *NullInt) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt{}

		in.Skip()

		return
	}

	*n = NullInt{Int: in.Int(), Valid: true}
}
