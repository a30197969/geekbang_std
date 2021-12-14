package main

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	done := make(chan error, 2)
	fmt.Println(cap(done))
}
