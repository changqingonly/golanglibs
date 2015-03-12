package golang_collections

import (
	"testing"
	"fmt"
)

func Test2(t *testing.T) {
	var uf UnionFind
	uf.Init(8)
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)
	uf.Union(6, 7)
	uf.Union(0, 2)
	uf.Union(4, 6)
	uf.Union(0, 4)
	uf.Dump()
	uf.Flatten()
	uf.Dump()
}
func Test1(t *testing.T) {
	var uf UnionFind
	uf.Init(8)
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)
	uf.Union(6, 7)
	uf.Union(0, 2)
	uf.Union(4, 6)
	uf.Union(0, 4)
	c := make(chan int)

	uf.Dump()
	go func() {
		uf.Flatten()
		c <- 0
	}()
	go func() {
		if uf.GroupCount() != 1 {
			t.Error("error")
		}
		if uf.Connected(0, 7) != true {
			t.Error("error")
		}
		c <- 0
	}()
	<-c
	<-c
	if uf.GroupCount() != 1 {
		t.Error("error")
	}
	if uf.Connected(0, 7) != true {
		t.Error("error")
	}
	uf.Dump()
}

func init() {
	fmt.Println("")
}