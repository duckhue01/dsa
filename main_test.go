package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		_ = sum(1, 10)
	})
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
