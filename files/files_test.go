package golang_files

import (
	"fmt"
	"reflect"
	"testing"
)

/*
go test -v -run=^TestReadLines2$
 */
func TestReadLines2(t *testing.T) {
	c := ReadZipLines(`words.zip`)
	for line := range c {
		fmt.Println(line)
	}
}

/*
go test -v -run=^TestReadLines1$
 */
func TestReadLines1(t *testing.T) {
	lines := []string{"aabdcd", "倒数"}

	for _, line := range lines {
		a := line
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(a, "<<<")
	}
}
