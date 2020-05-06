package nel

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type pointersSmall struct {
	field000 *string
	field001 *string
	field002 *string
	field003 *string
	field004 *string
	field005 *string
}

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
		field000: pointerSmallString(),
		field001: pointerSmallString(),
		field002: pointerSmallString(),
		field003: pointerSmallString(),
		field004: pointerSmallString(),
		field005: pointerSmallString(),
	}
}

func callPointerSmall(v *pointersSmall) {
}

func callPointerSmallJSON(v []byte) {
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
		callPointerSmallJSON(enc)
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
