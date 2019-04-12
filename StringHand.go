package Common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const leeters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//生成固定长度的随机字符串
func Urandom(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = leeters[rand.Intn(len(leeters))]
	}
	return b
}

//返回16进制的随机串
func GetSaltStr(n int) string {
	return hex.EncodeToString(Urandom(n))
}

//对用户名和密码加密
func NamePaswdToHash256(name, paswd, salt string) string {
	fmt.Println(salt)
	data := name + paswd + salt
	sum := sha256.Sum256([]byte(data))
	sumStr := fmt.Sprintf("%x", sum)
	ret := string(sumStr) + salt
	return ret
}
