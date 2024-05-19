package main

import (
	"log"
	"testing"

	"github.com/jzes/entendendo-algoritmos.git/toolBox"
)

func TestBinaryS(t *testing.T) {
	list := toolBox.MakeRange(100, 600)

	for i, v := range list {
		position, err := binaryS(list, v)
		if err != nil {
			log.Fatal(err)
		}
		if position != i {
			t.Errorf("want %d, got %d", i, position)
		}
	}
}

func BenchmarkBinaryS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryS(toolBox.MakeRange(100, 600), 300)
	}
}
