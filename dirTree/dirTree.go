package dirTree

/*
import (
	//"dirtreeHash/SHA1"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ostype = os.Getenv("GOOS") // 获取系统类型  windows or linux
)
var Listfile []string = nil

//  文件列表

func listfiles(path string, f os.FileInfo, err error) error {
	var strRet string // 路劲转义字符
	strRet, _ = os.Getwd()
	if ostype == "windows" {
		strRet += "\\"
	} else if ostype == "linux" {
		strRet += "/"
	}
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	strRet += path
	ok := strings.HasSuffix(strRet, suffix) // 若不是文件夹，则将路径加到Listfile里
	//ok1 :=strings.HasSuffix(strRet, ".txt")  // 过滤.txt结尾的文件
	//ok1 :=strings.HasSuffix(strRet, ".html") // 过滤.html结尾的文件
	//ok1 :=strings.HasSuffix(strRet, ".docx")  // 过滤.docx结尾的文件
	//ok1 :=strings.HasSuffix(strRet, ".ini")
	//Listfile = make([]string, 10240)
	fmt.Println("OK :", ok)
	if ok {
		Listfile = append(Listfile, strRet)
		//GetSha1(strRet)
	}
	return nil
}
func getFileList(path string) string {
	err := filepath.Walk(path, listfiles)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		panic(err)
	}
	//Listfile = Listfile[:len(Listfile)]
	return " "
}
func PrintFileList(path string) []string {
	getFileList(path)
	//Listfile = Listfile[:len(Listfile)]
	fmt.Println("文件路径1：", Listfile)
	return Listfile
}

*/
