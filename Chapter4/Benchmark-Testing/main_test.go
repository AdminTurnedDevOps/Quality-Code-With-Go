package main

import (
	"log"
	"testing"
)

func BenchmarkAddition(b *testing.B) {
	if b == nil {
		log.Println("The testing package contains no values")
	}

	got := addition(4)

	if got != 4 {
		b.Error("Did not equal 4")
	}
}
