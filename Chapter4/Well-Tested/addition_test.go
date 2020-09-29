package main

import (
	"log"
	"testing"
)

func TestAddition(t *testing.T) {
	if t == nil {
		log.Println("The testing package contains no values")
	}

	got := addition(2)

	if got != 4 {
		t.Error("Did not equal 4")
	}
}
