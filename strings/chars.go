package golang_strings

func ToCharArray(text string) []rune {
	return []rune(text)
}

func ClearRunes(cs []rune, initV rune) {
	for i, _ := range cs {
		cs[i] = initV
	}
}

// find common prefix
func FindCommonPrefix(remEntry, shortcut []rune) (commonPrefix []rune, x int) {
	smallestLen := len(remEntry)
	if smallestLen > len(shortcut) {
		smallestLen = len(shortcut)
	}
	for x = 0; x < smallestLen && shortcut[x] == remEntry[x]; x++ {

	}
	commonPrefix = shortcut[0:x]
	return
}

