package nan

import (
	"bytes"
	"fmt"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

func init() {
	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullTime",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			val := iter.ReadString()
			if val == "" {
				iter.ReportError("NullTime", "invalid value for time decoding \"\"")
				return
			}

			tm, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				msg := fmt.Sprintf("%v: invalid value for time decoding \"%s\"", err, val)
				iter.ReportError("NullTime", msg)
				return
			}

			*((*NullTime)(ptr)) = NullTime{Time: tm, Valid: true}
		},
	)

	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullTime",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullTime)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteString(t.Time.Format(time.RFC3339Nano))
		},
		nil,
	)
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return jsoniter.Marshal(n.Time.Format(time.RFC3339Nano))
}

func (n *NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*n = NullTime{}
		return nil
	}

	var dec string

	err := jsoniter.Unmarshal(data, &dec)
	if err != nil {
		return err
	}

	res, err := time.Parse(time.RFC3339Nano, string(dec))
	if err != nil {
		return err
	}

	*n = NullTime{Time: res, Valid: true}

	return nil
}

func (n NullTime) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.String(n.Time.Format(time.RFC3339Nano))
}

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
