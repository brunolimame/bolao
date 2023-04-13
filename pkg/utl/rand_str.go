package utl

import (
	"math/rand"
	"time"
)

func RandStr(length int) string {
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+=[]{}|;':\",.<>?/`~")
	if length <= 0 {
		length = 1
	}
	rand.Seed(time.Now().UnixNano()) // skipcq: GO-S1033

	b := make([]rune, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))] // skipcq: GSC-G404
	}
	return string(b)
}
