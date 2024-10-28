package main

import (
	"log"
	"testing"
)

func TestIntegretion(t *testing.T) {
	domath1 := doMath(2, 4, add)
	domath2 := doMath(5, 2, subtract)

	if domath1 != 6 || domath2 != 3 {
		log.Fatalf("There is something wrong!")
	}
}
