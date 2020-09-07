package nan

import (
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// MarshalEasyJSON - marshaller for easyjson
func (n nullTemplateType) MarshalEasyJSON(out *jwriter.Writer) {
	if !n.Valid {
		out.RawString("null")
		return
	}

	out.
	// out.Bool(n.Bool)
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *nullTemplateType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = nullTemplateType{}

		in.Skip()

		return
	}

	// *n = NullBool{Bool: in.Bool(), Valid: true}
}
