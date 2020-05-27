package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("main")
	m.Run()
}

func TestPrint(t *testing.T) {
	// t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("wrong")
	}
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}
