package double

import (
	"testing"
)

func BenchmarkDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Double(1)
	}
}
