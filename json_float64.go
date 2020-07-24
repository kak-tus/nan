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
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullFloat64)(ptr)) = NullFloat64{Float64: iter.ReadFloat64(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullFloat64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteFloat64(t.Float64)
		},
		nil,
	)
}

func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Float64)
}

func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat64{}
		return nil
	}

	var res float64

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullFloat64{Float64: res, Valid: true}

	return nil
}

func (n NullFloat64) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Float64(n.Float64)
}

func (n *NullFloat64) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullFloat64{}

		in.Skip()

		return
	}

	*n = NullFloat64{Float64: in.Float64(), Valid: true}
}
