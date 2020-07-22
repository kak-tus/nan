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
		"nan.NullBool",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullBool)(ptr)) = NullBool{Bool: iter.ReadBool(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullBool",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullBool)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteBool(t.Bool)
		},
		nil,
	)
}

func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Bool)
}

func (n *NullBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullBool{}
		return nil
	}

	var res bool

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullBool{Bool: res, Valid: true}

	return nil
}

func (n NullBool) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Bool(n.Bool)
}

func (n *NullBool) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullBool{}

		in.Skip()

		return
	}

	*n = NullBool{Bool: in.Bool(), Valid: true}
}
