package golang_strings

import (
"fmt"
)

const bmdebug = false

func preBmBc(ps []rune) (bmBc map[rune]int) {
	bmBc = make(map[rune]int)
	//刚匹配第一个字就出现坏字符, 则需要从模式串中找坏字符,
	//那么, 只能从下标是len(ps)-2的位置开始找
	for i := 0; i < len(ps)-1; i++ {
		bmBc[ps[i]] = len(ps)-1-i
	}
	return
}

func preBmGs(ps []rune) (bmGs []int) {
	m := len(ps)
	bmGs = make([]int, m)
	suff := suffix(ps)

	// case3
	for i := 0; i < m; i++ {
		bmGs[i] = m
	}

	//case2
	var j int
	for i := m - 1; i >= 0; i-- {
		if (suff[i] == i+1) {
			for ; j < m-1-i; j++ {
				if bmGs[j] == m {
					bmGs[j] = m-1-i
				}
			}
		}
	}

	// case1
	for i := 0; i <= len(ps)-2; i++ {
		bmGs[m-1-suff[i]] = m-1-i
	}
	return
}

func suffix(ps []rune) (suff []int) {
	m := len(ps)
	suff = make([]int, m)
	var i, j int
	suff[m-1] = m
	for i = m-2; i > -1; i-- {
		for j = i; j >= 0 && ps[j] == ps[m-1-(i-j)]; j-- { }
		suff[i] = i-j
	}
	return
}

func moveBmBc(t rune, ps []rune, j int, bmBc map[rune]int) int {
	v, ok := bmBc[t]
	if !ok {
		v = len(ps)
	}
	m := v - len(ps) + 1 + j
	if m <= 0 {
		return 1
	}
	return m
}

func BoyerMoore(text, pattern string) (pos int) {
	return BoyerMooreWithStart(text, pattern, 0)
}

func BoyerMooreWithStart(text, pattern string, start int) (pos int) {
	pos = -1
	ts, ps := []rune(text), []rune(pattern)
	bmBc := preBmBc(ps)
	bmGs := preBmGs(ps)
	for i := start; i <= len(ts)-len(ps); i++ {
		for j := len(ps) - 1; j > -1; j-- {
			if ps[j] == ts[i+j] {
				if j == 0 {
					pos = i
					return
				}
			}else {
				i+=max(moveBmBc(ts[i+j], ps, j, bmBc), bmGs[j])-1
				if bmdebug {
					fmt.Println("i:=", i)
				}
				break
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}







