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

func newPointerSmall() *pointersSmall {
	return &pointersSmall{
		Field000: pointerSmallString(),
		Field001: pointerSmallString(),
		Field002: pointerSmallString(),
		Field003: pointerSmallString(),
		Field004: pointerSmallString(),
		Field005: pointerSmallString(),
	}
}

func callPointerSmall(v *pointersSmall) {
}

func callPointerSmallJSON(v []byte) *pointersSmall {
	var res *pointersSmall
	_ = jsoniter.Unmarshal(v, res)

	return res
}

func callPointerSmallEasyJSON(v []byte) *pointersSmall {
	res := &pointersSmall{}
	_ = res.UnmarshalJSON(v)

	return res
}

func callPointerSmallA(v *pointersSmall) {
	callPointerSmallA1(v)
	callPointerSmallA2(v)
	callPointerSmallA3(v)
}

func callPointerSmallA1(v *pointersSmall) {
}

func callPointerSmallA2(v *pointersSmall) {
}

func callPointerSmallA3(v *pointersSmall) {
}

func callPointerSmallB(v *pointersSmall) {
	callPointerSmallB1(v)
	callPointerSmallB2(v)
	callPointerSmallB3(v)
}

func callPointerSmallB1(v *pointersSmall) {
}

func callPointerSmallB2(v *pointersSmall) {
}

func callPointerSmallB3(v *pointersSmall) {
}

func callPointerSmallC(v *pointersSmall) {
	callPointerSmallC1(v)
	callPointerSmallC2(v)
	callPointerSmallC3(v)
}

func callPointerSmallC1(v *pointersSmall) {
}

func callPointerSmallC2(v *pointersSmall) {
}

func callPointerSmallC3(v *pointersSmall) {
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
		dec := callPointerSmallJSON(enc)
		callPointerSmall(dec)
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
		dec := callPointerSmallEasyJSON(enc)
		callPointerSmall(dec)
	}

	b.StopTimer()
}
