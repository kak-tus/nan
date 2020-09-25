package nan

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoderFunc(
		"nan.nullTemplateType",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			var val initialTemplateType
			iter.ReadVal(&val)

			*((*nullTemplateType)(ptr)) = nullTemplateType{NullTemplateValue: val, Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.nullTemplateType",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			n := *((*nullTemplateType)(ptr))

			if !n.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteVal(n.NullTemplateValue)
		},
		nil,
	)
}
