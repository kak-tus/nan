package nan

import (
	"testing"

	"github.com/kak-tus/nan"
)

func BenchmarkAddrNan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addr := nan.Bool(true).Addr()
		_ = addr
	}
}

func BenchmarkAddrNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := true
		addr := &val
		_ = addr
	}
}
