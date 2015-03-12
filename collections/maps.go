package golang_collections

func RemoveDuplicationRune(a []rune) (b []rune) {
	b = make([]rune, 0, len(a))
	m := make(map[rune]interface{})
	for _, item := range a {
		m[item] = nil
	}
	for k, _ := range m {
		b = append(b, k)
	}
	return
}