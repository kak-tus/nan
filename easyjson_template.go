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

	n.NullTemplateValue.MarshalEasyJSON(out) // EasyJSON template
	// JSON template
	enc, err := n.NullTemplateValue.MarshalJSON() // JSON template
	if err != nil {                               // JSON template
		out.Error = err // JSON template
		return          // JSON template
	} // JSON template
	// JSON template
	out.RawString(string(enc)) // JSON template
}

// UnmarshalEasyJSON - unmarshaller for easyjson
func (n *nullTemplateType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		*n = nullTemplateType{}

		in.Skip()

		return
	}

	var val initialTemplateType

	val.UnmarshalEasyJSON(in) // EasyJSON template

	err := val.UnmarshalJSON(in.Raw()) // JSON template
	if err != nil {                    // JSON template
		in.AddError(err) // JSON template
		return           // JSON template
	} // JSON template

	*n = nullTemplateType{NullTemplateValue: val, Valid: true}
}
