package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("init")
	res := m.Run()
	fmt.Println("done");
	os.Exit(res)
}
