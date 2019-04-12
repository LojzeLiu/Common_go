package Common

import (
	"os"
)

type FileHand struct {
}

func GetFileHand() *FileHand {
	tmp := &FileHand{}
	return tmp
}

/*
功能：判断文件是否存在
参数：path 文件路径
返回：true：文件存在，false：文件不存在
*/
func (this *FileHand) FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/*
功能：判断文件是否为目录
参数：path 文件路径
返回：true：路径为目录，false：路径不为目录
*/
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/*
功能：判断文件是否为文件
参数：path 文件路径
返回：true：路径为文件，false：路径为目录
*/
func IsFile(path string) bool {
	return !IsDir(path)
}
