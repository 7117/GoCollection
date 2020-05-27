package main

import (
	"testing"
	"fmt"

)

func TestMain(m *testing.M){
	fmt.Println("main");
	m.Run();
}

func TestPrint(t *testing.T){
	// t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("wrong")
	}
}