package golang_collections

import (
	"fmt"
)

const debug = false

// 进化史: See page 222 (quick-find),page 224 (quick-union) andpage 228 (weighted).
type UnionFind struct {
	nodes           []int
	nodeWeights     []int
	nodeGroups      int
}

func (u *UnionFind) GroupCount() int {
	return u.nodeGroups
}

func (u *UnionFind) Init(nodeCount int) {
	u.nodeGroups = nodeCount
	u.nodes = make([]int, nodeCount, nodeCount)
	u.nodeWeights = make([]int, nodeCount, nodeCount)
	for i := 0; i < nodeCount; i++ {
		u.nodes[i] = i
		u.nodeWeights[i] = 1
	}
}

func (u *UnionFind) find(a int) int {
	originalA := a
	for a != u.nodes[a] {
		a = u.nodes[a]
	}
	u.nodes[originalA] = a
	return a
}

func (u *UnionFind) Union(a, b int) {
	ga := u.find(a)
	gb := u.find(b)
	if ga == gb {
		return
	}
	if u.nodeWeights[ga] < u.nodeWeights[gb] {
		// 保持ga比较大
		ga, gb = gb, ga
	}
	u.nodes[gb] = ga
	u.nodeWeights[ga]+=u.nodeWeights[gb]
	u.nodeGroups--
}

func (u *UnionFind) Connected(a, b int) bool {
	return u.find(a) == u.find(b)
}

func (u *UnionFind) Flatten() {
	for i := 0; i < len(u.nodes); i++ {
		u.find(i)
	}
}

func (u *UnionFind) Dump() {
	if debug {
		fmt.Println(u.nodes)
	}
}


