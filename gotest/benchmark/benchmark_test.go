package benchmark

import (
	"testing"
)

//go test -bench=.
//go test -bench=. -benchmem
//基准测试
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, aplane, a canal:Panma")
	}
}
