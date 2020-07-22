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
		"nan.NullInt64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt64)(ptr)) = NullInt64{Int64: iter.ReadInt64(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt64(t.Int64)
		},
		nil,
	)
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Int64)
}

func (n *NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullInt64{}
		return nil
	}

	var res int64

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullInt64{Int64: res, Valid: true}

	return nil
}

func (n NullInt64) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Int64(n.Int64)
}

func (n *NullInt64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullInt64{}

		in.Skip()

		return
	}

	*n = NullInt64{Int64: in.Int64(), Valid: true}
}
