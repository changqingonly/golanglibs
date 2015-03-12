package golang_strings

import (
	"testing"
	"fmt"
	"reflect"
)

// go test -v -run=^Test3$
func Test3(t *testing.T) {
	a := "啊的时代爱的asd多少事的"
	for _, bb := range []rune(a) {
		fmt.Printf("%#U\n", bb)
	}
}

// go test -v -run=^Test2$
func Test2(t *testing.T) {
	a := make([]rune, 0, 0);
	a = append(a, 332)
	a = append(a, 332)
	a = append(a, 332)
	fmt.Println(len(a), cap(a))
}

// go test -v -run=^Test1$
func Test1(t *testing.T) {
	var a  [3]int
	a[0] = 3
	fmt.Println(len(a), cap(a))
	b := a[:]
	fmt.Println(append(b, 33))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
