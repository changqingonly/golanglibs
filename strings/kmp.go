package golang_strings

import (
	"fmt"
	"reflect"
)

//const debug = true
const debug = false

func KMPFind(text string, key string) (pos int) {
	pos = KMPFindWithStart(0, text, key)
	return
}

// KMP 字符串匹配算法, 根据一下文章开发:
// http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
func KMPFindWithStart(start int, text string, key string) (pos int) {
	k := []rune(key)
	j := calcJump(k);
	if debug {
		fmt.Println(j)
	}
	t := []rune(text)
	pos = -1
	i, ii := start, 0
	for ; i < len(t); i++ {
		for ; ii < len(k) && i+ii < len(t); ii++ {
			if t[i+ii] == k[ii] {
				if ii == len(k)-1 {
					pos = i
					return
				}
			}else {
				if ii-1 >= 0 {
					i+=ii-j[ii-1]-1//主串前进, 移动位数 = 已匹配的字符数 - 对应的部分匹配值
					ii = j[ii-1]//模式串, 不必从0开始.
				}else {
					ii = 0
				}
				break
			}
		}
	}
	return
}

func calcJump(key []rune) (jump []int) {
	l := len(key)
	jump = make([]int, l, l)
	jump[0] = 0
	for i := 1; i < l; i++ {
		p := calcPrefix(key[:i + 1])
		s := calcSuffix(key[:i + 1])
		if debug {
			fmt.Println(key[:i + 1])
			fmt.Println(p)
			fmt.Println(s)
			fmt.Println("----")
		}
		jump[i] = commLen(p, s)
	}
	return
}

func calcPrefix(key []rune) (p [][]rune) {
	p = make([][]rune, 0)
	for i := 1; i < len(key); i++ {
		p = append(p, key[0:i])
	}
	return
}
func calcSuffix(key []rune) (s [][]rune) {
	s = make([][]rune, 0)
	for i := len(key) - 1; i > 0; i-- {
		s = append(s, key[i:len(key)])
	}
	return
}
func commLen(p, s [][]rune) (l int) {
	for i := 0; i < len(p); i++ {
		if reflect.DeepEqual(p[i], s[i]) {
			l = len(p[i])
			break
		}
	}
	return
}

func init() {
	fmt.Println("")
}