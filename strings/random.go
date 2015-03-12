package golang_strings

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuwxyzABCDEFGHIJKLMNOPQRSTUWXYZ0123456789"

var alphabetRunes = []rune(alphabet)
var alphabetRunesLen = len(alphabetRunes)

func RandomStr(length int) string {
	rand.Seed(int64(time.Now().UnixNano()))
	bs := make([]rune, length)
	for i := 0 ; i < length; i++ {
		bs[i] = alphabetRunes[rand.Intn(alphabetRunesLen)]
	}
	return string(bs)
}

