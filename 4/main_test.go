package main

import (
	"log"
	"testing"
)

func assert(b bool) {
	if !b {
		log.Fatal("Failed")
	}
}

func notassert(b bool) {
	assert(!b)
}

func TestPal(t *testing.T) {
	assert(isPal(9009))
	assert(isPal(101))
	notassert(isPal(50))
	notassert(isPal(100))
	notassert(isPal(102))
}
