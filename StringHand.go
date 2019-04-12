package Common

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const leeters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//生成固定长度的随机字符串
func Urandom(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = leeters[rand.Intn(len(leeters))]
	}
	return string(b)
}
