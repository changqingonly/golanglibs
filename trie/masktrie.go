package golang_trie

func generateMask(chars []rune) int {
	mask := len(chars)
RETRY:
	set := make(map[rune]interface{})
	for i, c := range chars {
		set[c&rune(mask)] = nil
		if len(set) != i+1 {
			mask++
			goto RETRY
		}
	}
	return mask
}

