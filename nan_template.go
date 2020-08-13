package nan

// nullTemplateType - nullable type for templates generation
type nullTemplateType struct {
	Value initialTemplateType
	Valid bool // Valid is true if Value is not NULL
}

func (n nullTemplateType) IsValid() bool {
	return n.Valid
}
