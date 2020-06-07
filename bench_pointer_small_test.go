package nan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func pointerSmallString() *string {
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

	return &s
}

func newPointerSmall() *pointerSmall {
	return &pointerSmall{
		Field000: pointerSmallString(),
		Field001: pointerSmallString(),
		Field002: pointerSmallString(),
		Field003: pointerSmallString(),
		Field004: pointerSmallString(),
		Field005: pointerSmallString(),
	}
}

func callPointerSmall(v *pointerSmall) {
}

func callPointerSmallJSON(v *pointerSmallJSON) {
}

func decodePointerSmallJSON(v []byte) *pointerSmallJSON {
	res := &pointerSmallJSON{}
	_ = jsoniter.Unmarshal(v, res)

	return res
}

func decodePointerSmallEasyJSON(v []byte) *pointerSmall {
	res := &pointerSmall{}
	_ = res.UnmarshalJSON(v)

	return res
}

func callPointerSmallA(v *pointerSmall) {
	callPointerSmallA1(v)
	callPointerSmallA2(v)
	callPointerSmallA3(v)
}

func callPointerSmallA1(v *pointerSmall) {
}

func callPointerSmallA2(v *pointerSmall) {
}

func callPointerSmallA3(v *pointerSmall) {
}

func callPointerSmallB(v *pointerSmall) {
	callPointerSmallB1(v)
	callPointerSmallB2(v)
	callPointerSmallB3(v)
}

func callPointerSmallB1(v *pointerSmall) {
}

func callPointerSmallB2(v *pointerSmall) {
}

func callPointerSmallB3(v *pointerSmall) {
}

func callPointerSmallC(v *pointerSmall) {
	callPointerSmallC1(v)
	callPointerSmallC2(v)
	callPointerSmallC3(v)
}

func callPointerSmallC1(v *pointerSmall) {
}

func callPointerSmallC2(v *pointerSmall) {
}

func callPointerSmallC3(v *pointerSmall) {
}

func BenchmarkPointersSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newPointerSmall()
		callPointerSmall(v)
	}

	b.StopTimer()
}

func BenchmarkPointersSmallJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newPointerSmall()

		enc, _ := jsoniter.Marshal(v)
		dec := decodePointerSmallJSON(enc)
		callPointerSmallJSON(dec)
	}

	b.StopTimer()
}

func BenchmarkPointerSmallChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newPointerSmall()
		callPointerSmallA(v)
		callPointerSmallB(v)
		callPointerSmallC(v)
	}

	b.StopTimer()
}

func BenchmarkPointersSmallEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newPointerSmall()

		enc, _ := v.MarshalJSON()
		dec := decodePointerSmallEasyJSON(enc)
		callPointerSmall(dec)
	}

	b.StopTimer()
}
