package snowflake

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	fmt.Printf("node id: %b\n", generateNodeID())
	fmt.Printf("timestamp: %b\n", ts())
	fmt.Printf("full token: %b\n", Next())
	// t.Fail()
}

func BenchmarkNext(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Next()
	}
}

func BenchmarkNextParallel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go Next()
	}
}
