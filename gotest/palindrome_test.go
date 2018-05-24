package gotest

import (
	"testing"
	"fmt"
)

//go test -v
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

func TestNonPalindromecn(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"Able was I ere I saw Elba", true},
		{"desserts", false},
		{"张三张", true},
		{"张思", false},
	}

	for _, test := range tests {
		if got := IsPalindromecn(test.input); got != test.want {
			t.Error("IsPalindromecn(%q) =%v", test.input, got)
		}
	}
}
