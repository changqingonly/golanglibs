package golang_trie

import "testing"

func TestMultipleAdditions(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("", "0")
	val, subpath := trie.GetWord("")
	if val.(string) != "0" {
		t.Errorf("Failed to retrieve entry with empty key")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("booboo", "1")
	val, subpath = trie.GetWord("booboo")
	if val.(string) != "1" {
		t.Errorf("Failed to retrieve first entry")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("boogoo", "2")
	val, subpath = trie.GetWord("boogoo")
	if val.(string) != "2" {
		t.Errorf("Failed to retrieve entry after mid split")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("boodoo", "3")
	val, subpath = trie.GetWord("boodoo")
	if val.(string) != "3" {
		t.Errorf("Failed to retrieve entry after additional mid split")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("boodod", "4")
	val, subpath = trie.GetWord("boodod")
	if val.(string) != "4" {
		t.Errorf("Failed to retrieve entry after tail variation")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("aoodod", "5")
	val, subpath = trie.GetWord("aoodod")
	if val.(string) != "5" {
		t.Errorf("Failed to retrieve entry after lead variation")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	trie.AddWord("你好世界", "6")
	val, subpath = trie.GetWord("你好世界")
	if val.(string) != "6" {
		t.Errorf("Failed to retrieve unicode entry")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}

	val, subpath = trie.GetWord("")
	if val.(string) != "0" {
		t.Errorf("Second sweep: Failed to retrieve entry with empty key")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("booboo")
	if val.(string) != "1" {
		t.Errorf("Second sweep: Failed to retrieve first entry")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("boogoo")
	if val.(string) != "2" {
		t.Errorf("Second sweep: Failed to retrieve entry after mid split")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("boodoo")
	if val.(string) != "3" {
		t.Errorf("Second sweep: Failed to retrieve entry after additional mid split")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("boodod")
	if val.(string) != "4" {
		t.Errorf("Second sweep: Failed to retrieve entry after tail variation")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("aoodod")
	if val.(string) != "5" {
		t.Errorf("Second sweep: Failed to retrieve entry after lead variation")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	val, subpath = trie.GetWord("你好世界")
	if val.(string) != "6" {
		t.Errorf("Second sweep: Failed to retrieve unicode entry")
	}
	if subpath != true {
		t.Errorf("Valid response with invalid subpath")
	}
	//	trie.DumpTree()
}

func TestValidPaths(t *testing.T) {
	trie := NewTrie()

	trie.AddWord("aaaaa", "1")
	val, validPath := trie.GetWord("aaa")
	if val != nil {
		t.Errorf("Value returned for subpath")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}
	trie.AddWord("aab", "2")
	val, validPath = trie.GetWord("aaa")
	if val != nil {
		t.Errorf("Value returned for subpath " + val.(string))
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}
	trie.AddWord("aaba", "3")
	val, validPath = trie.GetWord("aaa")
	if val != nil {
		t.Errorf("Value returned for subpath " + val.(string))
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}
	trie.AddWord("abaa", "4")
	val, validPath = trie.GetWord("abb")
	if val != nil {
		t.Errorf("Value returned for subpath " + val.(string))
	}
	if validPath == true {
		t.Errorf("Valid false positive")
	}

}

func TestMins(t *testing.T) {
	trie := NewTrie()

	trie.AddWord("a", "1")
	val, validPath := trie.GetWord("a")
	if val == nil {
		t.Errorf("Couldn't get a")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	trie.AddWord("b", "2")
	val, validPath = trie.GetWord("b")
	if val == nil {
		t.Errorf("Couldn't get b")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	trie.AddWord("aa", "3")
	val, validPath = trie.GetWord("aa")
	if val == nil {
		t.Errorf("Couldn't get aa")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}
}

func TestStuff(t *testing.T) {
	trie := NewTrie()

	trie.AddWord("ebay", "1")
	val, validPath := trie.GetWord("ebay")
	if val.(string) != "1" {
		t.Errorf("Unable to retrieve ebay 1")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	trie.AddWord("ebays", "2")
	val, validPath = trie.GetWord("ebays")
	if val.(string) != "2" {
		t.Errorf("Unable to retrieve ebays 2")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	trie.AddWord("eba", "4")
	trie.AddWord("ebay asdf", "5")

	val, validPath = trie.GetWord("ebay")
	if val.(string) != "1" {
		t.Errorf("Unable to retrieve ebay 3")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

}

func TestValidBranching(t *testing.T) {

	trie := NewTrie()
	trie.AddWord("shure asdf", "7")
	trie.AddWord("shure qwer", "8")
	trie.AddWord("shurtrax max-pax", "9")
	trie.AddWord("shura no toki", "10")
	trie.AddWord("shure", "6")

	val, validPath := trie.GetWord("shure")
	if val.(string) != "6" {
		t.Errorf("Unable to retrieve shure 6")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	val, validPath = trie.GetWord("shure asdf")
	if val.(string) != "7" {
		t.Errorf("Unable to retrieve shure 7")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	val, validPath = trie.GetWord("shure qwer")
	if val.(string) != "8" {
		t.Errorf("Unable to retrieve shure 8")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	val, validPath = trie.GetWord("shurtrax max-pax")
	if val.(string) != "9" {
		t.Errorf("Unable to retrieve shure 9")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

	val, validPath = trie.GetWord("shura no toki")
	if val.(string) != "10" {
		t.Errorf("Unable to retrieve shure 10")
	}
	if validPath != true {
		t.Errorf("Valid subpath not identified")
	}

}

func TestExteme(t *testing.T) {
	trie := NewTrie()

	trie.AddWord("+", "1")
	trie.AddWord(",", "1")
	trie.AddWord(".", "1")
	trie.AddWord(".a", "1")
	trie.AddWord("..", "1")
	trie.AddWord("[", "1")
	trie.GetWord("[")
	trie.GetWord(".")
	trie.GetWord("..")
	trie.GetWord(".+")
	trie.GetWord("[")
}
