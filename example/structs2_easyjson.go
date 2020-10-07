// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package example

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2e0fcc88DecodeGithubComKakTusNanExample(in *jlexer.Lexer, out *MyStruct2) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Val1":
			(out.Val1).UnmarshalEasyJSON(in)
		case "Val2":
			(out.Val2).UnmarshalEasyJSON(in)
		case "Val3":
			(out.Val3).UnmarshalEasyJSON(in)
		case "Val4":
			(out.Val4).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2e0fcc88EncodeGithubComKakTusNanExample(out *jwriter.Writer, in MyStruct2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Val1\":"
		out.RawString(prefix[1:])
		(in.Val1).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Val2\":"
		out.RawString(prefix)
		(in.Val2).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Val3\":"
		out.RawString(prefix)
		(in.Val3).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Val4\":"
		out.RawString(prefix)
		(in.Val4).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MyStruct2) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2e0fcc88EncodeGithubComKakTusNanExample(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MyStruct2) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2e0fcc88EncodeGithubComKakTusNanExample(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MyStruct2) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2e0fcc88DecodeGithubComKakTusNanExample(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MyStruct2) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2e0fcc88DecodeGithubComKakTusNanExample(l, v)
}
