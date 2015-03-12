package golang_trie

import "testing"
import "fmt"
import "os"
import "container/list"
import "strings"
import "encoding/binary"

/*
一般树结构的 trie 没 map 快, 但trie擅长多key查找
 */

var phraseFile string = `words.txt`

func BenchmarkInsertTrie(b *testing.B) {
	b.StopTimer()
	trie := NewTrie()
	f, err := os.Open(phraseFile)
	keys := list.New()
	for x := 0; x < b.N; x++ {
		if err != nil {
			f.Seek(0, 0)
		}
		var length uint32
		err = binary.Read(f, binary.LittleEndian, &length)
		str := make([]byte, length)
		_, err = f.Read(str)
		var gc uint32
		err = binary.Read(f, binary.LittleEndian, &gc)
		sstr := strings.ToUpper(string(str))
		sstr += fmt.Sprintf("%d", x)
		keys.PushBack(sstr)
	}
	b.StartTimer()
	x := 0
	for e := keys.Front(); e != nil; e = e.Next() {
		x++
		trie.AddWord(e.Value.(string), x)
	}
}

func BenchmarkInsertHash(b *testing.B) {
	b.StopTimer()
	f, err := os.Open(phraseFile)
	keys := list.New()
	for x := 0; x < b.N; x++ {
		if err != nil {
			f.Seek(0, 0)
		}
		var length uint32
		err = binary.Read(f, binary.LittleEndian, &length)
		str := make([]byte, length)
		_, err = f.Read(str)
		var gc uint32
		err = binary.Read(f, binary.LittleEndian, &gc)
		sstr := strings.ToUpper(string(str))
		sstr += fmt.Sprintf("%d", x)
		keys.PushBack(sstr)
	}
	b.StartTimer()
	var x uint32 = 0
	values := make(map[string]uint32, b.N)
	for e := keys.Front(); e != nil; e = e.Next() {
		x++
		values[e.Value.(string)] = x
	}
}

func BenchmarkFetchTrie(b *testing.B) {
	b.StopTimer()
	trie := NewTrie()
	f, err := os.Open(phraseFile)
	keys := list.New()
	for x := 0; x < b.N; x++ {
		if err != nil {
			f.Seek(0, 0)
		}
		var length uint32
		err = binary.Read(f, binary.LittleEndian, &length)
		str := make([]byte, length)
		_, err = f.Read(str)
		var gc uint32
		err = binary.Read(f, binary.LittleEndian, &gc)
		sstr := strings.ToUpper(string(str))
		sstr += fmt.Sprintf("%d", x)
		trie.AddWord(sstr, gc)
		keys.PushBack(sstr)
	}

	b.StartTimer()
	for e := keys.Front(); e != nil; e = e.Next() {
		_, _ = trie.GetWord(e.Value.(string))
	}
}

func BenchmarkFetchHashmap(b *testing.B) {
	b.StopTimer()
	f, err := os.Open(phraseFile)
	keys := list.New()
	values := make(map[string]uint32, b.N)
	for x := 0; x < b.N; x++ {
		if err != nil {
			f.Seek(0, 0)
		}
		var length uint32
		err = binary.Read(f, binary.LittleEndian, &length)
		str := make([]byte, length)
		_, err = f.Read(str)
		var gc uint32
		err = binary.Read(f, binary.LittleEndian, &gc)
		sstr := strings.ToUpper(string(str))
		sstr += fmt.Sprintf("%d", x)
		values[sstr] = gc
		keys.PushBack(sstr)
	}

	b.StartTimer()
	for e := keys.Front(); e != nil; e = e.Next() {
		_ = values[e.Value.(string)]
	}
}

