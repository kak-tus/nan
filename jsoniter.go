package nan

import (
	"fmt"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullBool",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullBool)(ptr)) = NullBool{Bool: iter.ReadBool(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullBool",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullBool)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteBool(t.Bool)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullBool)(ptr)).Valid
		},
	)

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
		func(ptr unsafe.Pointer) bool {
			return !((*NullFloat32)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullFloat64)(ptr)) = NullFloat64{Float64: iter.ReadFloat64(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullFloat64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullFloat64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteFloat64(t.Float64)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullFloat64)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt)(ptr)) = NullInt{Int: iter.ReadInt(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt(t.Int)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullInt)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt8",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt8)(ptr)) = NullInt8{Int8: iter.ReadInt8(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt8",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt8)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt8(t.Int8)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullInt8)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt16",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt16)(ptr)) = NullInt16{Int16: iter.ReadInt16(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt16",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt16)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt16(t.Int16)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullInt16)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt32)(ptr)) = NullInt32{Int32: iter.ReadInt32(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt32(t.Int32)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullInt32)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullInt64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullInt64)(ptr)) = NullInt64{Int64: iter.ReadInt64(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullInt64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullInt64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteInt64(t.Int64)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullInt64)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullString)(ptr)) = NullString{String: iter.ReadString(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullString",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullString)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteString(t.String)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullString)(ptr)).Valid
		},
	)

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
		func(ptr unsafe.Pointer) bool {
			return !((*NullTime)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullUint",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullUint)(ptr)) = NullUint{Uint: iter.ReadUint(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullUint",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullUint)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteUint(t.Uint)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullUint)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullUint8",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullUint8)(ptr)) = NullUint8{Uint8: iter.ReadUint8(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullUint8",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullUint8)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteUint8(t.Uint8)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullUint8)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullUint16",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullUint16)(ptr)) = NullUint16{Uint16: iter.ReadUint16(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullUint16",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullUint16)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteUint16(t.Uint16)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullUint16)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullUint32",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullUint32)(ptr)) = NullUint32{Uint32: iter.ReadUint32(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullUint32",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullUint32)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteUint32(t.Uint32)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullUint32)(ptr)).Valid
		},
	)

	jsoniter.RegisterTypeDecoderFunc(
		"nan.NullUint64",
		func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
			if iter.ReadNil() {
				return
			}

			*((*NullUint64)(ptr)) = NullUint64{Uint64: iter.ReadUint64(), Valid: true}
		},
	)
	jsoniter.RegisterTypeEncoderFunc(
		"nan.NullUint64",
		func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
			t := *((*NullUint64)(ptr))

			if !t.Valid {
				stream.WriteNil()
				return
			}

			stream.WriteUint64(t.Uint64)
		},
		func(ptr unsafe.Pointer) bool {
			return !((*NullUint64)(ptr)).Valid
		},
	)

}
