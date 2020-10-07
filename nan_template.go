package nan

// nullTemplateType - generated for nullTemplateType
//easyjson:skip
type nullTemplateType struct {
	NullTemplateValue initialTemplateType
	Valid             bool // Valid is true if Value is not NULL
}

func naninitialTemplateType(v initialTemplateType) nullTemplateType {
	return nullTemplateType{NullTemplateValue: v, Valid: true}
}

func (n nullTemplateType) IsValid() bool {
	return n.Valid
}

func (n nullTemplateType) IsDefined() bool {
	return n.Valid
}
