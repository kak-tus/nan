package nan

// nullTemplateType - generated for nullTemplateType
type nullTemplateType struct {
	Value initialTemplateType
	Valid bool // Valid is true if Value is not NULL
}

func (n nullTemplateType) IsValid() bool {
	return n.Valid
}
