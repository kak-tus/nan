package nel

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type valuesSmall struct {
	field000 string
	field001 string
	field002 string
	field003 string
	field004 string
	field005 string
}

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

func newValueSmall() valuesSmall {
	return valuesSmall{
		field000: valueSmallString(),
		field001: valueSmallString(),
		field002: valueSmallString(),
		field003: valueSmallString(),
		field004: valueSmallString(),
		field005: valueSmallString(),
	}
}

func callValueSmall(v valuesSmall) {
}

func callValueSmallJSON(v []byte) {
}

func callValueSmallA(v valuesSmall) {
	callValueSmallA1(v)
	callValueSmallA2(v)
	callValueSmallA3(v)
}

func callValueSmallA1(v valuesSmall) {
}

func callValueSmallA2(v valuesSmall) {
}

func callValueSmallA3(v valuesSmall) {
}

func callValueSmallB(v valuesSmall) {
	callValueSmallB1(v)
	callValueSmallB2(v)
	callValueSmallB3(v)
}

func callValueSmallB1(v valuesSmall) {
}

func callValueSmallB2(v valuesSmall) {
}

func callValueSmallB3(v valuesSmall) {
}

func callValueSmallC(v valuesSmall) {
	callValueSmallC1(v)
	callValueSmallC2(v)
	callValueSmallC3(v)
}

func callValueSmallC1(v valuesSmall) {
}

func callValueSmallC2(v valuesSmall) {
}

func callValueSmallC3(v valuesSmall) {
}

func BenchmarkValuesSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()
		callValueSmall(v)
	}

	b.StopTimer()
}

func BenchmarkValuesSmallJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()

		enc, _ := jsoniter.Marshal(v)
		callValueSmallJSON(enc)
	}

	b.StopTimer()
}

func BenchmarkValuesSmallChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newValueSmall()
		callValueSmallA(v)
		callValueSmallB(v)
		callValueSmallC(v)
	}

	b.StopTimer()
}
