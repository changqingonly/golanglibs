package golang_trie

import (
	"container/list"
)

type acTrieState struct {
	success map[rune]*acTrieState
	failure *acTrieState
	emit []rune
}

func createAcTrieState() (s *acTrieState) {
	s = new(acTrieState)
	s.success = make(map[rune]*acTrieState)
	s.emit = nil
	s.failure = nil
	return
}

func (s*acTrieState) addRune(c rune) *acTrieState {
	ns, ok := s.success[c]
	if !ok {
		ns = createAcTrieState()
		s.success[c] = ns
	}
	return ns
}

type AcTrie struct {
	rootState *acTrieState
}

/*
	如果同时命中, abba,与aa, 则, aa被忽略.
	如果有, a1,aa1,da1会同时命中
 */
func CreateAcTrie() (t*AcTrie) {
	t = new(AcTrie)
	t.rootState = createAcTrieState()
	return
}

func (t*AcTrie) AddKeyword(k string) {
	kw := []rune(k)
	state := t.rootState
	for _, c := range kw {
		state = state.addRune(c)
	}
	state.emit = kw
}

func (t*AcTrie) BuildAcTrie() {
	queue := list.New()
	t.rootState.failure = t.rootState// 根失败边为根

	// 深度为1的节点, failure 边为 root
	for _, v := range t.rootState.success {
		v.failure = t.rootState
		queue.PushBack(v)
	}

	// 对于深度为 >1 的节点, 做广度优先搜索
	for queue.Len() > 0 {
		currentState := queue.Front().Value.(*acTrieState)
		queue.Remove(queue.Front())

		for c, subState := range currentState.success {
			queue.PushBack(subState)

			// currentState 经过 c 到达 subState
			// 找出subState的failure边
			subState.failure = currentState.failure.getNextState(c)
		}
	}
}

func (state*acTrieState) getNextState(c rune) (next*acTrieState) {
	var ok bool
	next, ok = state.success[c]
	for ; !ok; next, ok = state.success[c] {
		if state == state.failure {// root state
			next = state
			break
		}
		state = state.failure
	}
	return
}

func (t*AcTrie) AcTrieFindKeys(text string) (keys []string, cleanText string) {
	mkeys := make(map[string]interface{})
	ts := []rune(text)
	clean := make([]rune, len(ts))
	state := t.rootState
	for i, c := range ts {
		clean[i] = c
		state = state.getNextState(c)
		if len(state.emit) > 0 {
			mkeys[string(state.emit)] = nil
			for j := 0; j < len(state.emit); j++ {
				clean[i-len(state.emit)+1+j] = '*'
			}
		}
	}
	cleanText = string(clean)

	keys = make([]string, 0, len(mkeys))
	for k, _ := range mkeys {
		keys = append(keys, k)
	}
	return
}


