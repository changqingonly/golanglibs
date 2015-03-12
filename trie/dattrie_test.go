package golang_trie

import (
	"testing"
	"fmt"
	"github.com/renrenxing/golanglibs/files"
)

func TestDAT9(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	lines := golang_files.ReadZipLines(zipfile)
	for line := range lines {
		myTrie.Insert(line)
	}
	myTrie.BuildDATTrie()
	keys, cleanText := myTrie.FindKeys("af十打傻了都较为服务点多坐台法萨看到是恩爱维持额阿尔法啊哒哒额")
	//	for _, k := range keys {
	//		fmt.Println(k)
	//	}
	//	fmt.Println(cleanText)
	if len(keys) != 1 {
		t.Error("错误")
	}
	if cleanText.String() != "af十打傻了都较为服务点多*法萨看到是恩爱维持额阿尔法啊哒哒额" {
		t.Error("错误")
	}
	myTrie.Statistic()
}
func TestDAT8_1(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	myTrie.Insert("啊哒哒额啊哒哒额啊哒哒额啊哒哒额")
	myTrie.Insert("啊哒哒额啊哒哒额啊哒哒额")
	myTrie.Insert("啊哒哒额啊哒哒额")
	myTrie.Insert("啊哒哒额")
	myTrie.Insert("的口味的的")
	myTrie.Insert("打傻了都较为服务")
	myTrie.Insert("萨看到是恩爱维持")
	myTrie.BuildDATTrie()
	keys, cleanText := myTrie.FindKeys("asdfdasf啊哒哒额啊哒哒额啊哒哒额啊哒哒额dasfasdf")
	//	for _, k := range keys {
	//		fmt.Println(k)
	//	}
	//	fmt.Println(cleanText)
	if len(keys) != 1 {
		t.Error("错误")
	}
	if cleanText.String() != "asdfdasf*dasfasdf" {
		t.Error("错误")
	}
}
func TestDAT8(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	myTrie.Insert("啊哒哒额")
	myTrie.Insert("的口味的的")
	myTrie.Insert("打傻了都较为服务")
	myTrie.Insert("萨看到是恩爱维持")
	myTrie.BuildDATTrie()
	keys, cleanText := myTrie.FindKeys("啊哒哒额a啊哒哒额啊哒哒额da的口味的的saf十打傻了都较为服务点多法萨看到是恩爱维持额阿尔法啊哒哒额")
	if len(keys) != 7 {
		t.Error("错误")
	}
	if cleanText.String() != "*a**da*saf十*点多法*额阿尔法*" {
		t.Error("错误")
	}
}
func TestDAT7(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	myTrie.Insert("abcdef32哈哈32呵呵2dsdsddsdsdsdsd")
	myTrie.Insert("abcdef32322dsdsddsd")
	myTrie.Insert("斯大林的322")
	myTrie.Insert("abcdef3")
	myTrie.Insert("abc鄂温克def")
	myTrie.Insert("aebcdef")
	myTrie.Insert("12aebcdef")
	myTrie.Insert("ffbcdef")
	myTrie.Insert("3323223")
	myTrie.Insert("332的 到时3223")
	myTrie.BuildDATTrie()
	_, cleanText := myTrie.FindKeys("d12aebcdefsacabcdef32哈哈32呵呵2dsdsddsdsdsdsdc斯大林的3223斯大林的322")
	if cleanText.String() != "d*sac*c*3*" {
		t.Error("错误")
	}
}
func TestDAT6(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	myTrie.Insert("abcdef32322dsdsddsdsdsdsd")
	myTrie.Insert("abcdef32322dsdsddsd")
	myTrie.Insert("abcdef32322")
	myTrie.Insert("abcdef3")
	myTrie.Insert("abcdef")
	myTrie.Insert("aebcdef")
	myTrie.Insert("12aebcdef")
	myTrie.Insert("ffbcdef")
	myTrie.Insert("3323223")
	myTrie.Insert("3323223")
	myTrie.BuildDATTrie()
	_, cleanText := myTrie.FindKeys("3323223adsdffbcdefsabcdef32322dsdsddsdd12aebcdefsdsdsdds3323223")
	if cleanText.String() != "*adsd*s*d*sdsdsdds*" {
		t.Error("错误")
	}
}
func TestDAT5(t *testing.T) {
	myTrie := new(DATTrie)
	myTrie.Init()
	myTrie.Insert("abcdef32322dsdsddsdsdsdsd")
	myTrie.Insert("abcdef32322dsdsddsd")
	myTrie.Insert("abcdef32322")
	myTrie.Insert("abcdef3")
	myTrie.Insert("abcdef")
	myTrie.Insert("aebcdef")
	myTrie.Insert("12aebcdef")
	myTrie.Insert("ffbcdef")
	myTrie.Insert("3323223")
	myTrie.Insert("3323223")
	myTrie.BuildDATTrie()

	xx := myTrie.SearchPrefixes([]rune("abcdef32322dsdsddsdsdsdsd"))
	if len(xx) != 5 {
		t.Error("错误")
	}
}
func TestDAT4(t *testing.T) {
	trie := new(DATTrie)
	trie.Init()
	trie.Insert("s")
	trie.BuildDATTrie()
	{
		x := trie.SearchMaxPrefix([]rune("abcdef32322"))
		if len(x) != 0 {
			t.Error("错误")
		}
	}
	{
		x := trie.SearchMaxPrefix([]rune("s"))
		if string(x) != "s" {
			t.Error("错误")
		}
	}
}

func TestDAT3(t *testing.T) {
	trie := new(DATTrie)
	trie.Init()
	trie.Insert("")
	trie.BuildDATTrie()
	x := trie.SearchMaxPrefix([]rune("abcdef32322"))
	if len(x) != 0 {
		t.Error("错误")
	}
}

func TestDAT2(t *testing.T) {
	trie := new(DATTrie)
	trie.Init()
	trie.BuildDATTrie()
	x := trie.SearchMaxPrefix([]rune("abcdef32322"))
	if len(x) != 0 {
		t.Error("错误")
	}
}

func TestDAT1(t *testing.T) {
	trie := new(DATTrie)
	trie.Init()
	trie.Insert("abcdef32322")
	trie.Insert("abcdef3")
	trie.Insert("abcdef")
	trie.Insert("aebcdef")
	trie.Insert("12aebcdef")
	trie.BuildDATTrie()

	{
		x := trie.SearchMaxPrefix([]rune("abcdef32322"))
		if string(x) != "abcdef32322" {
			t.Error("错误")
		}
	}
	{
		x := trie.SearchMaxPrefix([]rune("abcdsdef32322"))
		if len(x) != 0 {
			t.Error("错误")
		}
	}
}

func init() {
	fmt.Println()
}

var zipfile string = `words.zip`
