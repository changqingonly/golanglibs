package golang_strings

import (
	"testing"
	"reflect"
	"fmt"
)

func TestKMP7(t *testing.T) {
	text := "qwf2fqw32qff阿尔法ew2efsdfa阿尔法"
	s := "阿尔法"
	p := KMPFindWithStart(13, text, s)
	if p != 24 {
		t.Error("错误")
	}
}
func TestKMP6(t *testing.T) {
	text := "qwf2fqw32qff阿尔法ew2efsdfa阿尔法"
	s := "阿尔法"
	p := KMPFindWithStart(12, text, s)
	if p != 12 {
		t.Error("错误")
	}
}
func TestKMP5(t *testing.T) {
	text := "qwf2fqw32qff阿尔法"
	s := "阿尔法"
	p := KMPFind(text, s)
	if p != 12 {
		t.Error("错误")
	}
}
func TestKMP4(t *testing.T) {
	text := "阿尔法啊的的发发阿尔法啊发点啥法师法额阿尔法"
	s := "阿尔法"
	p := KMPFind(text, s)
	if p != 0 {
		t.Error("错误")
	}
}
func TestKMP3(t *testing.T) {
	text := "啊的的发发阿尔法啊发点啥法师法额"
	s := "阿尔法"
	p := KMPFind(text, s)
	if p != 5 {
		t.Error("错误")
	}
}
func TestKMP2(t *testing.T) {
	text := "BBC ABCDAB ABCDABCDABDE"
	s := "ABCDABD"
	p := KMPFind(text, s)
	if p != 15 {
		t.Error("错误")
	}
}

func TestKMP1(t *testing.T) {
	s := "abcdabcd"
	j := calcJump([]rune(s))
	if reflect.DeepEqual(j, []int{0, 0, 0, 0, 1, 2, 3, 4}) != true {
		t.Error("错误")
	}
}

func init() {
	fmt.Println("")
}