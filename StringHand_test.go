package Common

import (
	"fmt"
	"testing"
)

//func GetSaltStr(n int) string {
func TestGetSaltStr(t *testing.T) {
	salt := GetSaltStr(8)
	if len(salt) != 16 {
		t.Error("GetSaltStr(failed), salt:", salt)
	}
	fmt.Println("salt:", salt, ";")
}

//func NamePaswdToHash256(name, paswd, salt string) string
func TestNamePaswdToHash256(t *testing.T) {
	sum := NamePaswdToHash256("lojze", "abc78520", GetSaltStr(8))
	fmt.Println(sum)
}
