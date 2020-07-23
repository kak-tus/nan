package nan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func valueSmallString() string {
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

func newValueSmall() valueSmall {
	return valueSmall{
		Field000: valueSmallString(),
		Field001: valueSmallString(),
		Field002: valueSmallString(),
		Field003: valueSmallString(),
		Field004: valueSmallString(),
		Field005: valueSmallString(),
	}
}

func newValueSmallJSON() valueSmallJSON {
	return valueSmallJSON{
		Field000: valueSmallString(),
		Field001: valueSmallString(),
		Field002: valueSmallString(),
		Field003: valueSmallString(),
		Field004: valueSmallString(),
		Field005: valueSmallString(),
	}
}

func callValueSmall(v valueSmall) {
}

func callValueSmallJSON(v valueSmallJSON) {
}

func decodeValueSmallJSON(v []byte) valueSmallJSON {
	var res valueSmallJSON
	_ = jsoniter.Unmarshal(v, &res)

	return res
}

func decodeValueSmallEasyJSON(v []byte) valueSmall {
	var res valueSmall
	_ = res.UnmarshalJSON(v)

	return res
}

func callValueSmallA(v valueSmall) {
	callValueSmallA1(v)
	callValueSmallA2(v)
	callValueSmallA3(v)
}

func callValueSmallA1(v valueSmall) {
}

func callValueSmallA2(v valueSmall) {
}

func callValueSmallA3(v valueSmall) {
}

func callValueSmallB(v valueSmall) {
	callValueSmallB1(v)
	callValueSmallB2(v)
	callValueSmallB3(v)
}

func callValueSmallB1(v valueSmall) {
}

func callValueSmallB2(v valueSmall) {
}

func callValueSmallB3(v valueSmall) {
}

func callValueSmallC(v valueSmall) {
	callValueSmallC1(v)
	callValueSmallC2(v)
	callValueSmallC3(v)
}

func callValueSmallC1(v valueSmall) {
}

func callValueSmallC2(v valueSmall) {
}

func callValueSmallC3(v valueSmall) {
}

func BenchmarkValueSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()
		callValueSmall(v)
	}
}

func BenchmarkValueSmallJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmallJSON()

		enc, _ := jsoniter.Marshal(v)
		dec := decodeValueSmallJSON(enc)
		callValueSmallJSON(dec)
	}
}

func BenchmarkValueSmallChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()
		callValueSmallA(v)
		callValueSmallB(v)
		callValueSmallC(v)
	}
}

func BenchmarkValueSmallEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()

		enc, _ := v.MarshalJSON()
		dec := decodeValueSmallEasyJSON(enc)
		callValueSmall(dec)
	}
}
