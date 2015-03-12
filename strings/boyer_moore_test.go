package golang_strings

import (
"testing"
"fmt"
"strings"
)

func Test_BoyerMoore4(t *testing.T) {
	s := "HERE IS A SIMPLE EXAMPLE"
	sep := "EXAMPLE"
	pos := BoyerMoore(s, sep)
	pos2 := strings.Index(s, sep)
	if pos != pos2 {
		t.Error("err")
	}
}

func Test_BoyerMoore5(t *testing.T) {
	s := "aaaaaaaaaaaaaaaabaaaaaaaaaaa"
	sep := "aaabaaaa"
	pos := BoyerMoore(s, sep)
	pos2 := strings.Index(s, sep)
	if pos != pos2 {
		t.Error("err")
	}
}

func Test_BoyerMoore3(t *testing.T) {
	s := "mahtavaatalomaisema omalomailuun"
	sep := "maisemaomaloma"
	pos := BoyerMoore(s, sep)
	pos2 := strings.Index(s, sep)
	if pos != pos2 {
		t.Error("err")
	}
}

func Test_BoyerMoore2(t *testing.T) {
	s := "abacccb bbazz"
	sep := "cbadccb"
	pos := BoyerMoore(s, sep)
	pos2 := strings.Index(s, sep)
	if pos != pos2 {
		t.Error("err")
	}
}

func Test_BoyerMoore1(t *testing.T) {
	s := "fjaksdfj;efsdfase"
	sep := "fas"
	pos := BoyerMoore(s, sep)
	pos2 := strings.Index(s, sep)
	if pos != pos2 {
		t.Error("err")
	}
}

func init() {
	fmt.Println("")
}
