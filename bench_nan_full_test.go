package nan

import (
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func newNanFull() nanFull {
	return nanFull{
		Field000: StringToNullString("7676"),
		Field001: BoolToNullBool(true),
		Field002: Int64ToNullInt64(7676),
		Field003: TimeToNullTime(time.Now()),
	}
}

func newNanFullJSON() nanFullJSON {
	return nanFullJSON{
		Field000: StringToNullString("7676"),
		Field001: BoolToNullBool(true),
		Field002: Int64ToNullInt64(7676),
		Field003: TimeToNullTime(time.Now()),
	}
}

func callNanFull(v nanFull) {
}

func callNanFullJSON(v nanFullJSON) {
}

func decodeNanFullJSON(v []byte) nanFullJSON {
	var res nanFullJSON
	_ = jsoniter.Unmarshal(v, &res)

	return res
}

func decodeNanFullEasyJSON(v []byte) nanFull {
	var res nanFull
	_ = res.UnmarshalJSON(v)

	return res
}

func callNanFullA(v nanFull) {
	callNanFullA1(v)
	callNanFullA2(v)
	callNanFullA3(v)
}

func callNanFullA1(v nanFull) {
}

func callNanFullA2(v nanFull) {
}

func callNanFullA3(v nanFull) {
}

func callNanFullB(v nanFull) {
	callNanFullB1(v)
	callNanFullB2(v)
	callNanFullB3(v)
}

func callNanFullB1(v nanFull) {
}

func callNanFullB2(v nanFull) {
}

func callNanFullB3(v nanFull) {
}

func callNanFullC(v nanFull) {
	callNanFullC1(v)
	callNanFullC2(v)
	callNanFullC3(v)
}

func callNanFullC1(v nanFull) {
}

func callNanFullC2(v nanFull) {
}

func callNanFullC3(v nanFull) {
}

func BenchmarkNanFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanFull()
		callNanFull(v)
	}

	b.StopTimer()
}

func BenchmarkNanFullChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanFull()
		callNanFullA(v)
		callNanFullB(v)
		callNanFullC(v)
	}

	b.StopTimer()
}

func BenchmarkNanFullJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanFullJSON()

		enc, _ := jsoniter.Marshal(v)
		dec := decodeNanFullJSON(enc)
		callNanFullJSON(dec)
	}

	b.StopTimer()
}

func BenchmarkNanFullEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := newNanFull()

		enc, _ := v.MarshalJSON()
		dec := decodeNanFullEasyJSON(enc)
		callNanFull(dec)
	}

	b.StopTimer()
}
