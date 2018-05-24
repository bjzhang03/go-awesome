package gotest

import (
	"testing"
	"fmt"
)

func TestIslindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error("IsPalindrome(\" detartrated \") = false")
	}

	if !IsPalindrome("kayak") {
		t.Error("IsPalindrome(\" kayak \") = false")
	}
	fmt.Println("Test Success!")
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palinderome") {
		t.Error("IsPalindrome(\" palinderome \") = true")
	}
}
