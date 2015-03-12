package golang_trie

import (
	"fmt"
	"strings"
	"math"
	"bytes"
	"time"
	"github.com/renrenxing/golang_strings"
)

//const debug = true
const debug = false
const dump = false
const nan = rune(0)

// 基于文档An Implementation of Double-Array Trie
// http://linux.thai.net/~thep/datrie/datrie.html#Double .
// golang 双数组的实现
// 但有所改变, 文档使用的是行初始化.
// 这里使用的是, 列初始化, 所以不需要复杂的 relocate .
// 应用1, 查询公共前缀 -> 可以用于敏感词过滤. 给一段 wordMaxLen 长度的字串, 给出有没有敏感词.
// 应用2, 给一个词, 查出以这个词开头的单词. 需要在每个DAT状态上附加一个单词集合(存放id),
// 		然后从一个大数组中根据idList取, 单词.
type DATTrie struct {
	base               []rune
	check              []rune
	baseArrLength      int               //定义base数组的长度
	wordMaxLen         int               //最长的单词字的个数
	offset             rune              //状态的偏移量
	wordCount          int               //单词个数
	startBuildTime     time.Time         //本trie何时被创建
	endBuildTime       time.Time         //本trie何时被创建完毕
	words         map[string]interface{} // 用于初始化, 初始化后, 清除
}

// 用于敏感词过滤
func (dat *DATTrie) FindKeys(text string) (keys []string, cleanText *bytes.Buffer) {
	keys = make([]string, 0)
	cleanText = new(bytes.Buffer)
	chars := []rune(text)
	textLen := len(chars)
	var p []rune
	for i := 0; i < textLen; i++ {
		p = dat.SearchMaxPrefix(chars[i:textLen])
		if len(p) > 0 {
			keys = append(keys, string(p))
			i += len(p)-1
			cleanText.WriteString("*")
		}else {
			cleanText.WriteRune(chars[i])
		}
	}
	return
}

// 查找 text 的最大前缀.
func (dat *DATTrie) SearchMaxPrefix(text []rune) (prefix []rune) {
	prefixes := dat.SearchPrefixes(text)
	if len(prefixes) > 0 {
		prefix = prefixes[len(prefixes)-1]
	}
	return
}

func (dat *DATTrie) SearchPrefixes(text []rune) (prefixes [][]rune) {
	prefixes = make([][]rune, 0)
	if len(text) > 0 {
		s := text[0]
		for i, c := range text {
			t, p, key := dat.walk(s, c)
			if !p {
				break
			}
			if key {
				prefixes = append(prefixes, text[:i+1])
			}
			s = t
		}
	}
	return
}

func (dat *DATTrie) walk(s , c rune) (t rune, validPath, isKey bool) {
	t = dat.getOffset(s)+c
	if state, key := dat.getCheck(t); state == s {
		validPath = true
		isKey = key
	}
	return
}

func (dat *DATTrie) InitWithBaseLen(baseLen int) {
	dat.baseArrLength = baseLen
	dat.base = make([]rune, dat.baseArrLength, dat.baseArrLength)
	dat.check = make([]rune, dat.baseArrLength, dat.baseArrLength)
	dat.words = make(map[string]interface{})
	dat.offset = 1
	dat.wordMaxLen = 0
}

func (dat *DATTrie) Init() {
	dat.InitWithBaseLen(10000 * 10)
}

func (dat *DATTrie) Insert(word string) {
	word = strings.Trim(word, "\n\r \t")
	if length := len([]rune(word)); length > 0 {
		if dat.wordMaxLen < length {
			dat.wordMaxLen = length
		}
		dat.words[word] = nil
	}
}

func (dat *DATTrie) BuildDATTrie() {
	defer func() { dat.endBuildTime = time.Now() }()
	defer func() {dat.words = nil/*释放内存, 不再使用map*/}()
	dat.startBuildTime = time.Now()
	dat.wordCount = len(dat.words)

	words := [][]rune{}
	for k, _ := range dat.words {
		words = append(words, []rune(k))
	}

	ss := make([]rune, dat.wordCount)
	cs := make([]rune, dat.wordCount)
	ts := make([]rune, dat.wordCount)
	for col := 0; col < dat.wordMaxLen+1/*加一, 给最长的敏感词设置结束符: check为负数*/; col++ {
		ss, ts = ts, ss
		for row := 0; row < dat.wordCount; row++ {
			cs[row] = nan
			if col < len(words[row]) {
				cs[row] = words[row][col]
			}
		}
		if col == 0 {
			dat.createS0(cs, ss)
		}
		if debug {
			fmt.Println("输入")
			fmt.Println(ss)
			fmt.Println(string(cs))
		}
		dat.transfer(ss, cs, ts)
		if debug {
			fmt.Println("输出")
			fmt.Println(ts)
			fmt.Println("")
		}
		if dump {
			dat.Dump()
		}
	}
}

// 每个状态 s 都有一个偏移量 offset, base[s]=offset
func (dat *DATTrie) generateOffset() rune {
	if int(dat.offset) == dat.baseArrLength/2-1 {
		dat.offset = 1
	}
	for {
		dat.offset++
		if dat.getOffset(dat.offset) == nan {
			return dat.offset
		}
	}
}

func (dat *DATTrie) getOffset(s rune) rune {
	return dat.base[s]
}

func (dat *DATTrie) setOffset(s, offset rune) {
	dat.base[s] = offset
}

func (dat *DATTrie) getCheck(t rune) (s rune, isKey bool) {
	tmp := float64(dat.check[t])
	isKey = tmp < 0
	s = rune(math.Abs(tmp))
	return
}

func (dat *DATTrie) setCheck(s, t rune, isKey bool) {
	x := 1
	if isKey {
		x = -1
	}
	dat.check[t] = rune(x)*s
}

// cs,ss 均为列向量
// 创建初始状态, 初始状态直接使用每个单词的第一个字
func (dat *DATTrie) createS0(cs, ss []rune) {
	for row, _ := range ss {
		ss[row] = cs[row]
		dat.setOffset(ss[row], 10086) // 先占后分配, 防止分配是冲突
	}
	for row, _ := range ss {
		dat.setOffset(ss[row], dat.generateOffset()) // 分配
	}
}

// ss,cs,ts 均为列向量
func (dat *DATTrie) transfer(ss, cs, ts []rune) {
RETRY:
	golang_strings.ClearRunes(ts, nan)
	for row := 0; row < len(ss); row++ {
		if cs[row] == nan {
			if ss[row] != nan {
				// 一个单词的结束
				pre_s, _ := dat.getCheck(ss[row])
				dat.setCheck(pre_s, ss[row], true)
			}
			continue
		}
		// 从 ss 经过 cs 字符到 ts 状态.
		ts[row] = dat.getOffset(ss[row])+cs[row]
		if dat.getOffset(ts[row]) == nan { // 没有被占用
			//占用
			dat.setOffset(ts[row], dat.generateOffset())
			//设置 check, 后续的trie单词查找要用到
			dat.setCheck(ss[row], ts[row], false)
		} else {
			pre_s , _ := dat.getCheck(ts[row])
			if pre_s != ss[row] {
				// 已经被占用, 需要 relocate
				dat.relocate(ss, cs, row)
				goto RETRY
			}
		}
	}
}

// 列初始化, 简单的 relocate
// 清除已经确定的状态转换,
func (dat *DATTrie) relocate(ss , cs []rune, row int) {
	s := ss[row] // 保存出问题的状态, 我们需要改变此状态的偏移量
	newOffset := dat.generateOffset()//生成新的偏移量
	for i := 0; i <= row; i++ {
		if s == ss[i] {//对于所有从这个状态出发的char
			if i != row {//对于已经做出状态转移的char, 要清空base与check
				t := dat.getOffset(s) + cs[i]
				dat.setOffset(t, nan)
				dat.setCheck(nan, t, false)// 清空check
			}
			dat.setOffset(ss[i], newOffset)//将出问题的状态的offset改变
		}
	}
}

//---------- debug 方法-----------
func (dat *DATTrie) StatisticStr() (out *bytes.Buffer) {
	out = new(bytes.Buffer)
	out.WriteString(fmt.Sprintf("双数组Trie于%s开始构建,构建耗时%v\n", dat.startBuildTime, dat.endBuildTime.Sub(dat.startBuildTime)))
	out.WriteString(fmt.Sprintf("一共有%d个敏感词\n", dat.wordCount))
	out.WriteString(fmt.Sprintf("最长的敏感词是有%d个字\n", dat.wordMaxLen))
	out.WriteString(fmt.Sprintf("base,check数组长度为%d\n", dat.baseArrLength))
	out.WriteString(fmt.Sprintf("base,check一共使用了%dKB的空间\n", dat.baseArrLength*32/8/1024*2))
	{
		var start, end int
		for i, c := range dat.base {
			if c != nan {
				start = i
				break
			}
		}
		for i := dat.baseArrLength - 1; i > -1; i-- {
			if dat.base[i] != nan {
				end = i
				break
			}
		}
		a := float64(start) / float64(dat.baseArrLength) * 100
		b := float64(dat.baseArrLength - end) / float64(dat.baseArrLength) * 100
		out.WriteString(fmt.Sprintf("base中:前面有%f%%的空间未使用,后面有%f%%的空间未使用\n", a, b))
	}
	return
}
func (dat *DATTrie) Statistic() {
	fmt.Println(dat.StatisticStr())
}
func (dat *DATTrie) Dump() {
	fmt.Println("base: ", dat.base)
	fmt.Println("check:", dat.check)
}


