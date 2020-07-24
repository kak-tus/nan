package nan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/kak-tus/nan"
)

func nanSmallString() string {
	s := "01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890" +
		"01234567890012345678900123456789001234567890012345678900123456789001234567890012345678900123456789001234567890"

	return s
}

func newNanSmall() nanSmall {
	return nanSmall{
		Field000: nan.StringToNullString(nanSmallString()),
		Field001: nan.StringToNullString(nanSmallString()),
		Field002: nan.StringToNullString(nanSmallString()),
		Field003: nan.StringToNullString(nanSmallString()),
		Field004: nan.StringToNullString(nanSmallString()),
		Field005: nan.StringToNullString(nanSmallString()),
	}
}

func newNanSmallJSON() nanSmallJSON {
	return nanSmallJSON{
		Field000: nan.StringToNullString(nanSmallString()),
		Field001: nan.StringToNullString(nanSmallString()),
		Field002: nan.StringToNullString(nanSmallString()),
		Field003: nan.StringToNullString(nanSmallString()),
		Field004: nan.StringToNullString(nanSmallString()),
		Field005: nan.StringToNullString(nanSmallString()),
	}
}

func callNanSmall(v nanSmall) {
}

func callNanSmallJSON(v nanSmallJSON) {
}

func decodeNanSmallJSON(v []byte) nanSmallJSON {
	var res nanSmallJSON
	_ = jsoniter.Unmarshal(v, &res)

	return res
}

func decodeNanSmallEasyJSON(v []byte) nanSmall {
	var res nanSmall
	_ = res.UnmarshalJSON(v)

	return res
}

func callNanSmallA(v nanSmall) {
	callNanSmallA1(v)
	callNanSmallA2(v)
	callNanSmallA3(v)
}

func callNanSmallA1(v nanSmall) {
}

func callNanSmallA2(v nanSmall) {
}

func callNanSmallA3(v nanSmall) {
}

func callNanSmallB(v nanSmall) {
	callNanSmallB1(v)
	callNanSmallB2(v)
	callNanSmallB3(v)
}

func callNanSmallB1(v nanSmall) {
}

func callNanSmallB2(v nanSmall) {
}

func callNanSmallB3(v nanSmall) {
}

func callNanSmallC(v nanSmall) {
	callNanSmallC1(v)
	callNanSmallC2(v)
	callNanSmallC3(v)
}

func callNanSmallC1(v nanSmall) {
}

func callNanSmallC2(v nanSmall) {
}

func callNanSmallC3(v nanSmall) {
}

func BenchmarkNanSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmall()
		callNanSmall(v)
	}
}

func BenchmarkNanSmallChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmall()
		callNanSmallA(v)
		callNanSmallB(v)
		callNanSmallC(v)
	}
}

func BenchmarkNanSmallJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmallJSON()

		enc, _ := jsoniter.Marshal(v)
		dec := decodeNanSmallJSON(enc)
		callNanSmallJSON(dec)
	}
}

func BenchmarkNanSmallEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmall()

		enc, _ := v.MarshalJSON()
		dec := decodeNanSmallEasyJSON(enc)
		callNanSmall(dec)
	}
}
