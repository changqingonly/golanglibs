package golang_trie

import (
	"fmt"
	"github.com/renrenxing/golanglibs/strings"
)

const startLetter = 32
const endLetter = 127

func (*Trie) createNodePtrArr() []*Node {
	numLetters := endLetter - startLetter + 1
	return make([]*Node, numLetters, numLetters)
}

// 获取未被分配的编码
func (this *Trie) nextCode() (code int) {
	code = this.code
	this.code++
	return code
}

// 返回某个字符的编码, 如果是中文, 则从map中取
// 每个char都有一个唯一标示
func (this *Trie) getCode(char rune, update bool) (code int) {
	code = -1
	if ir := int(char); ir >= startLetter && ir <= endLetter {
		// 字母的标示获取方式
		code = ir-startLetter
	} else {
		// 汉字的标示获取方式
		map_code, exists := this.unicodeMap[char]
		if !exists {
			if update {
				code = this.nextCode()
				this.unicodeMap[char] = code
			}
		} else {
			code = map_code
		}
	}
	return
}

// 扩展子节点数组
func (this *Trie) ensureCapacity(children []*Node, index int) []*Node {
	if len(children) < index+1 {
		for x := len(children); x < index+1; x++ {
			children = append(children, nil)
		}
	}
	return children
}

func (this *Trie) addNode(nd *Node, words []rune, value interface{}) (added bool) {
	//没有 common_prefix 则在Trie引用加上, 则这个节点是新的节点, 没有子节点
	//所以公共前缀就是 words
	if nd.common_prefix == nil {
		nd.common_prefix = words
		nd.value = value
		nd.children = nil
		added = true
		return
	}

	if len(words) == 0 && len(nd.common_prefix) == 0 { // 如果传入的是空串, 则直接替换 value, 递归的结束
		// 空串是所有的节点的公共节点
		nd.value = value
		return
	} else { // 找出最小公共前缀
		new_prefix, new_pref_len := golang_strings.FindCommonPrefix(words, nd.common_prefix)
		if new_pref_len < len(nd.common_prefix) { // 分裂当前节点, 出现更短的公共前缀, 需要断开, 一分二
			code := this.getCode(nd.common_prefix[new_pref_len], true)
			new_tail := nd.common_prefix[new_pref_len+1 : len(nd.common_prefix)] //common_prefix 非公共尾部
			newNode := &Node{
				children:      nd.children, //当前节点的子节点给, 新节点, 新节点作为当前节点的子节点
				value:         nd.value, //值
				common_prefix: new_tail,
			}
			nd.children = this.createNodePtrArr()//清空当前的子节点
			nd.children = this.ensureCapacity(nd.children, code)
			nd.children[code] = newNode
			nd.common_prefix = new_prefix
			nd.value = nil
		}
		if new_pref_len < len(words) { //递归, 添加子节点
			//尾巴第一个字, 保存在节点数组中, 剩下的递归到子节点.
			//公共前缀是表示压缩的多层. 正常来说, 每一个字保存在子节点数组中, 代表一个层次
			//只是使用公共前缀来压缩层次
			tail_first := this.getCode(words[new_pref_len], true)
			//尾部的余下部分
			tail_remain := words[new_pref_len+1 : len(words)]
			nd.children = this.ensureCapacity(nd.children, tail_first)
			if nd.children[tail_first] == nil {
				nd.children[tail_first] = &Node{
					children:      nil,
					value:         nil,
					common_prefix: nil,
				}
			}
			this.addNode(nd.children[tail_first], tail_remain, value)
		} else {//与公共前缀相等
			// 公共前缀就是 words
			nd.value = value
			added = true
		}
	}
	return
}

// 节点
// 1) 一个节点只存一个单词, 则单词保存在 common_prefix
//	a) children 为 nil
// 2) 如果一个节点有多个单词, 并且单词有公共前缀, 公共前缀保存在 common_prefix
//	a) 不同的尾巴中, 首字母保存在 children 上
// 3) 如果一个节点有多个单词, 但至少有2个单词的首字母不一样, 这 common_prefix 为空数组, 不为 nil
//	a) 不一样的首字母保存再 children 数组中
//	b) 后面的字母用同样的操作递归作用在子节点上
type Node struct {
	children      []*Node     //每个子节点下面有多个节点
	value         interface{} //每个有单词的节点都挂一个值
	common_prefix []rune      //共同的前缀则使用一个节点表示. 如果为空, 说明是结束节点
}

// 引用
type Trie struct {
	nd         *Node        //根节点
	unicodeMap map[rune]int //汉字以及对应的编码
	code       int          //不断递增, 为 unicodeMap 分配编码
	MaxLen     int
	KeyCount   int //todo 以后纠正
}

func (this *Trie) AddWord(word string, value interface{}) {
	if added := this.addNode(this.nd, []rune(word), value); added {
		this.KeyCount++
	}
}

func (this *Trie) GetWord(word string) (value interface{}, validPath bool) {
	curr_node := this.nd
	remain_chars := []rune(word)
	for char_offset := 0; char_offset <= len(remain_chars); char_offset++ {
		var comm_pref_offset int
		for comm_pref_offset = 0; comm_pref_offset < len(curr_node.common_prefix); comm_pref_offset++ {
			if char_offset+comm_pref_offset >= len(remain_chars) {
				//输入的单词相当于trie中单词的前缀, 无法匹配到, 查找失败
				return nil, true //返回true表示路径正确, 后面使用trie做多key匹配时有大用
			}
			if curr_node.common_prefix[comm_pref_offset] != remain_chars[char_offset+comm_pref_offset] {
				//公共前缀都不匹配则查找失败
				return nil, false //返回false表示路径错误
			}
		}
		// chars与当前节点的公共前缀的前一部分匹配, 则, char偏移量加上前缀的偏移量.
		char_offset += comm_pref_offset
		if char_offset < len(remain_chars) {
			// 到达此步说明前缀已经匹配成功, 接下来遇到分支, 需要做选择
			index := this.getCode(remain_chars[char_offset], false)
			if index < 0 || index > len(curr_node.children)-1 || curr_node.children[index] == nil {
				// 节点不存在, 查找失败
				return nil, false
			} else {
				// 匹配到了单词的前一部分, 需要继续匹配节点的子节点
				curr_node = curr_node.children[index]     //当前节点切换成子节点
				remain_chars = remain_chars[char_offset:] //前一部分已经匹配到, 只需要后一部分与子节点匹配
				char_offset = 0                           //切换到后一部分, 所以位移归零
			}
		}
	}
	return curr_node.value, true
}


func (this *Trie) IncLookup(nd *Node, prefOffset int, char rune) (value interface{}, validPath bool, nextNd*Node, nextOffset int) {
	if nd == nil {
		//使用根节点
		nd = this.nd
	}

	nextOffset = prefOffset
	nextNd = nd
	value = nd.value
	validPath = false

	travelDown := func() {
		defer func() {
			recover()
		}()
		if nd.children != nil {
			code := this.getCode(char, false)
			if nextNd = nd.children[code]; nextNd != nil {
				validPath = true
			}
		}
	}

	if len(nd.common_prefix) == 0 {
		travelDown()
	}else if nextOffset < len(nd.common_prefix) {
		if nd.common_prefix[nextOffset] == char {
			nextOffset++
			nextNd = nd
			validPath = true
		}else {
			validPath = false
		}
	} else {
		travelDown()
	}
	return
}

func NewTrie() *Trie {
	return &Trie{
		nd: &Node{
			children:      nil,
			value:         nil,
			common_prefix: nil,
		},
		unicodeMap: make(map[rune]int),
		code:       100, //中文从100起编码, 英文占用了前96个
	}
}

//------------------其他debug方法-----------

// 只支持dump英文单词
func (this *Trie) DumpTree() {
	this.dumpNode(this.nd, 1)
}

func indent(depth int, msg ...interface{}) {
	for x := 0; x < depth; x++ {
		fmt.Print("  ")
	}
	fmt.Println(fmt.Sprint(msg...))
}

func (this *Trie) dumpNode(t *Node, depth int) {
	indent(depth, "- commprefix: ", string(t.common_prefix))
	indent(depth, "- value: ", t.value)
	if t.children != nil {
		indent(depth, "- children:")
		for y := 0; y < len(t.children); y++ {
			if t.children[y] != nil {
				indent(depth, " - ", string(y + startLetter))
				this.dumpNode(t.children[y], depth+1)
			}
		}
	}
}
