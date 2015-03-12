package golang_trie

import (
	"testing"
	"fmt"
)

func ATestIncLookup(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("", "")
	trie.AddWord("", "")
	trie.AddWord("b", "b")
	trie.AddWord("1abcdef", "1abcdef")
	trie.AddWord("abcdef", "abcdef")
	trie.AddWord("abcdsef", "abcdsef")
	{
		val, subpath := trie.GetWord("abcdef")
		if !(val == "abcdef" && subpath) {
			t.Error("错误")
		}
	}
	{
		val, subpath := trie.GetWord("1abcdef")
		if !(val == "1abcdef" && subpath) {
			t.Error("错误")
		}
	}
//	trie.DumpTree()

	var n *Node
	var o int
	for _, char := range []rune("abcdefg") {
		v, p, n, o := trie.IncLookup(n, o, char)
		fmt.Println(v, p, n, o)
	}
}


func ATestAllLetter(*testing.T) {
	const startLetter = 32
	const endLetter = 127
	numLetters := int(endLetter) - int(startLetter) + 1
	fmt.Println(numLetters)
}



