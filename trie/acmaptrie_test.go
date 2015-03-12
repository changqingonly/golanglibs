package golang_trie

import (
	"testing"
	"fmt"
)

/*
go test -v -run=^TestACTrie5$
 */
func TestACTrie5(t *testing.T) {
	trie := CreateAcTrie()
	trie.AddKeyword("s吧d")
	trie.AddKeyword("啊s吧d吃")
	trie.AddKeyword("啊s吧d吃1")
	trie.AddKeyword("d啊s吧d吃")
	trie.AddKeyword("s吧d吃1")
	trie.BuildAcTrie()
	a, b := trie.AcTrieFindKeys("啊吧吃hea啊吧啊s吧d吃吃dfa啊吧吃fdhe啊吧吃aaa啊s吧d啊s吧d吃1啊s吧d吃aaa")
	//	fmt.Println(a, b)
	if len(a) != 4 {
//		t.Error("err")
	}
	if b != "啊吧吃hea啊吧*****吃dfa啊吧吃fdhe啊吧吃aaa啊s吧************aaa" {
		t.Error("err")
	}
}

func TestACTrie4(t *testing.T) {
	trie := CreateAcTrie()
	trie.AddKeyword("啊吧吃")
	trie.BuildAcTrie()
	a, b := trie.AcTrieFindKeys("啊吧吃hea啊吧吃dfa啊吧吃fdhe啊吧吃")
	if len(a) != 4 {
//		t.Error("err")
	}
	if b != "***hea***dfa***fdhe***" {
		t.Error("err")
	}
}

func TestACTrie3(t *testing.T) {
	trie := CreateAcTrie()
	trie.AddKeyword("啊吧吃")
	trie.BuildAcTrie()
	_, b := trie.AcTrieFindKeys("hea啊吧吃dfafdhe")
	if b != "hea***dfafdhe" {
		t.Error("err")
	}
}

func TestACTrie2(t *testing.T) {
	trie := CreateAcTrie()
	trie.AddKeyword("he")
	trie.BuildAcTrie()
	_, b := trie.AcTrieFindKeys("hehehehehe")
	if b != "**********" {
		t.Error("err")
	}
}

func TestACTrie1(t *testing.T) {
	trie := CreateAcTrie()
	trie.AddKeyword("hers")
	trie.AddKeyword("his")
	trie.AddKeyword("she")
	trie.AddKeyword("he")
	trie.BuildAcTrie()
	a, b := trie.AcTrieFindKeys("ushers")
	if len(a) != 2 {
//		t.Error("err")
	}
	if b != "u*****" {
		t.Error("err")
	}
}

func init() {
	fmt.Println()
}