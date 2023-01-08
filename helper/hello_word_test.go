package helper

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestHelloWord(t *testing.T) {
	result := HelloWord("Adhi")
	if result != "Hello Adhi" {
		t.Error("Result Must be 'Hello Adhi'")
	}
	fmt.Println("Hello Adhi Done")
}

func TestHelloW(t *testing.T) {
	result := HelloWord("Adhi")
	if result != "Hello Adhi" {
		t.Fatal("Result Must be 'Hello Adhi'")
	}
	fmt.Println("Hello Adhi Done")
}

func TestNum(t *testing.T) {
	var data bits = bits(127)
	fmt.Println(data)
}
