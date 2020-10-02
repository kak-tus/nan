package nan

import "github.com/gocql/gocql"

// MarshalCQL - marshaller for cql
func (n nullTemplateType) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.NullTemplateValue.MarshalCQL(info)
}

// UnmarshalCQL - unmarshaller for cql
func (n *nullTemplateType) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	if len(data) == 0 {
		*n = nullTemplateType{}
		return nil
	}

	var val initialTemplateType
	if err := val.UnmarshalCQL(info, data); err != nil {
		return err
	}

	*n = nullTemplateType{Valid: true, NullTemplateValue: val}

	return nil
}
