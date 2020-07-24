package nan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/kak-tus/nan"
)

func nanBigString() string {
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

func newNanBig() nanBig {
	return nanBig{
		Field000: nan.StringToNullString(nanBigString()),
		Field001: nan.StringToNullString(nanBigString()),
		Field002: nan.StringToNullString(nanBigString()),
		Field003: nan.StringToNullString(nanBigString()),
		Field004: nan.StringToNullString(nanBigString()),
		Field005: nan.StringToNullString(nanBigString()),
	}
}

func newNanBigJSON() nanBigJSON {
	return nanBigJSON{
		Field000: nan.StringToNullString(nanBigString()),
		Field001: nan.StringToNullString(nanBigString()),
		Field002: nan.StringToNullString(nanBigString()),
		Field003: nan.StringToNullString(nanBigString()),
		Field004: nan.StringToNullString(nanBigString()),
		Field005: nan.StringToNullString(nanBigString()),
	}
}

func callNanBig(v nanBig) {
}

func callNanBigJSON(v nanBigJSON) {
}

func decodeNanBigJSON(v []byte) nanBigJSON {
	var res nanBigJSON
	_ = jsoniter.Unmarshal(v, &res)

	return res
}

func decodeNanBigEasyJSON(v []byte) nanBig {
	var res nanBig
	_ = res.UnmarshalJSON(v)

	return res
}

func callNanBigA(v nanBig) {
	callNanBigA1(v)
	callNanBigA2(v)
	callNanBigA3(v)
}

func callNanBigA1(v nanBig) {
}

func callNanBigA2(v nanBig) {
}

func callNanBigA3(v nanBig) {
}

func callNanBigB(v nanBig) {
	callNanBigB1(v)
	callNanBigB2(v)
	callNanBigB3(v)
}

func callNanBigB1(v nanBig) {
}

func callNanBigB2(v nanBig) {
}

func callNanBigB3(v nanBig) {
}

func callNanBigC(v nanBig) {
	callNanBigC1(v)
	callNanBigC2(v)
	callNanBigC3(v)
}

func callNanBigC1(v nanBig) {
}

func callNanBigC2(v nanBig) {
}

func callNanBigC3(v nanBig) {
}

func BenchmarkNanBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanBig()
		callNanBig(v)
	}
}

func BenchmarkNanBigChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanBig()
		callNanBigA(v)
		callNanBigB(v)
		callNanBigC(v)
	}
}

func BenchmarkNanBigJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanBigJSON()

		enc, _ := jsoniter.Marshal(v)
		dec := decodeNanBigJSON(enc)
		callNanBigJSON(dec)
	}
}

func BenchmarkNanBigEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanBig()

		enc, _ := v.MarshalJSON()
		dec := decodeNanBigEasyJSON(enc)
		callNanBig(dec)
	}
}
