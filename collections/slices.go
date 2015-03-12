package golang_collections

func ContainsRune(a rune, arr[]rune) bool {
	for _, t := range arr {
		if a == t {
			return true
		}
	}
	return false
}

func Slice2MapRune(arr[]rune) map[rune]interface{} {
	m := make(map[rune]interface{})
	for i, t := range arr {
		m[t] = i
	}
	return m
}

/*

func SliceEq(a, b interface{}) bool {
	length := SliceGetLen(a)
	if length != SliceGetLen(b) {
		return false
	}
	for i := 0; i < length; i++ {
		if SliceGetN(a, i) != SliceGetN(b, i) {
			return false
		}
	}
	return true
}

func SliceGetN(slice interface{}, index int) (item interface{}) {
	switch t := slice.(type){
	default:
		panic("SliceGetN 只能处理数组, 且是原生数据类型, " + fmt.Sprintf("unexpected type %T", t))
	case []int:
		item = slice.([]int)[index]
	case []int16:
		item = slice.([]int16)[index]
	case []int32:
		item = slice.([]int32)[index]
	case []int64:
		item = slice.([]int64)[index]
	case []float32:
		item = slice.([]float32)[index]
	case []float64:
		item = slice.([]float64)[index]
	}
	return
}

func SliceGetLen(slice interface{}) (length int) {
	switch t := slice.(type){
	default:
		panic("SliceGetLen 只能处理数组, 且是原生数据类型, " + fmt.Sprintf("unexpected type %T", t))
	case []int:
		length = len(slice.([]int))
	case []int16:
		length = len(slice.([]int16))
	case []int32:
		length = len(slice.([]int32))
	case []int64:
		length = len(slice.([]int64))
	case []float32:
		length = len(slice.([]float32))
	case []float64:
		length = len(slice.([]float64))
	}
	return
}
 */

/*
package main

import (
	"unsafe"
)

func Equal(p1 unsafe.Pointer, s1 uintptr, n1 int, p2 unsafe.Pointer, s2 uintptr, n2 int) bool {
	if s1 != s2 || n1 != n2 {
		return false
	}
	n := s1 * uintptr(n1)
	for i := uintptr(0); i < n; i++ {
		b1 := (*byte)(p1)
		b2 := (*byte)(p2)
		if *b1 != *b2 {
			return false
		}
		p1 = unsafe.Pointer(uintptr(p1) + 1)
		p2 = unsafe.Pointer(uintptr(p2) + 1)
	}
	return true
}

func main() {
	type T struct {
		i int
		s string
	}
	var a = []T{{1, "a"}, {2, "b"}}
	var b = []T{{1, "a"}, {2, "b"}}
	if Equal(unsafe.Pointer(&a[0]), unsafe.Sizeof(a[0]), len(a), unsafe.Pointer(&b[0]), unsafe.Sizeof(b[0]), len(b)) {
		println("yes")
	} else {
		println("no")
	}
}

 */
