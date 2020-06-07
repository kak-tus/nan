package nan

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
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
		Field000: StringToNullString(nanSmallString()),
		Field001: StringToNullString(nanSmallString()),
		Field002: StringToNullString(nanSmallString()),
		Field003: StringToNullString(nanSmallString()),
		Field004: StringToNullString(nanSmallString()),
		Field005: StringToNullString(nanSmallString()),
	}
}

func callNanSmall(v nanSmall) {
}

func callNanSmallJSON(v []byte) nanSmall {
	var res nanSmall
	_ = jsoniter.Unmarshal(v, &res)

	return res
}

// func callNanSmallEasyJSON(v []byte) nanSmall {
// 	var res nanSmall
// 	_ = res.UnmarshalJSON(v)

// 	return res
// }

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

	b.StopTimer()
}

func BenchmarkNanSmallJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmall()

		enc, _ := jsoniter.Marshal(v)
		dec := callNanSmallJSON(enc)
		callNanSmall(dec)
	}

	b.StopTimer()
}

func BenchmarkNanSmallChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanSmall()
		callNanSmallA(v)
		callNanSmallB(v)
		callNanSmallC(v)
	}

	b.StopTimer()
}

// func BenchmarknanSmallEasyJSON(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		v := newNanSmall()

// 		enc, _ := v.MarshalJSON()
// 		dec := callNanSmallEasyJSON(enc)
// 		callNanSmall(dec)
// 	}

// 	b.StopTimer()
// }
