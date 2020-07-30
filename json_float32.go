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
		"nan.NullFloat32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullFloat32)(ptr)) = NullFloat32{Float32: iter.ReadFloat32(), Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullFloat32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullFloat32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteFloat32(t.Float32)
		},
		nil,
	)
}

func (n NullFloat32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Float32)
}

func (n *NullFloat32) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullFloat32{}
		return nil
	}

	var res float32

	err := jsoniter.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	*n = NullFloat32{Float32: res, Valid: true}

	return nil
}

func (n NullFloat32) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.Float32(n.Float32)
}

func (n *NullFloat32) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = NullFloat32{}

		in.Skip()

		return
	}

	*n = NullFloat32{Float32: in.Float32(), Valid: true}
}
