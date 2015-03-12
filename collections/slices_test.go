package golang_collections

import (
	"testing"
	"reflect"
)

type ss1 struct {
	a int
	p *ss1
}

func TestSlice5(t*testing.T) {
	{
		var a ss1
		var b ss1
		a.a = 3
		b.a = 3
		a.p = new(ss1)
		b.p = new(ss1)
		if !reflect.DeepEqual(a, b) {
			t.Error("err")
		}
	}
	{
		var a ss1
		var b ss1
		a.a = 3
		b.a = 3
		a.p = new(ss1)
		b.p = new(ss1)
		a.p.a = 3
		if reflect.DeepEqual(a, b) {
			t.Error("err")
		}
	}
}

func TestSlice4(t*testing.T) {
	{
		s1 := make([]rune, 10, 10)
		s2 := make([]rune, 10, 10)
		s1[3] = 44
		s2[3] = 44
		if !reflect.DeepEqual(s1, s2) {
			t.Error("err")
		}
	}
	{
		s1 := make([]rune, 110, 110)
		s2 := make([]rune, 10, 10)
		s1[3] = 44
		s2[3] = 44
		if reflect.DeepEqual(s1, s2) {
			t.Error("err")
		}
	}
	{
		s1 := make([]rune, 10, 10)
		s2 := make([]rune, 10, 10)
		s1[3] = 441
		s2[3] = 44
		if reflect.DeepEqual(s1, s2) {
			t.Error("err")
		}
	}
}

//
//func TestSlice3(t*testing.T) {
//	s1 := make([]rune, 10, 10)
//	s2 := make([]rune, 10, 10)
//	s1[3] = 44
//	s2[3] = 44
//	if SliceEq(s1, s2) != true {
//		t.Error("err")
//	}
//}
//
//func TestSlice2(t*testing.T) {
//	s := make([]rune, 10, 10)
//	s[3] = 44
//	if SliceGetN(s, 3).(rune) != 44 {
//		t.Error("err")
//	}
//}
//
//func TestSlice1(t*testing.T) {
//	fmt.Println()
//	s := make([]rune, 10, 10)
//	if SliceGetLen(s) != 10 {
//		t.Error("err")
//	}
//}
