package outFile

import (
	"dirtreeHash/SHA1"
	//"dirtreeHash/dirTree"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"
)

var OutputFileName string = "dirTreefile.txt"

//var listfile = GetfileList(Path)

// 计算文件大小
func checkErr(err error) {
	if err != nil {
		fmt.Println("Failed to create file... ", err)
		panic(err)
	}
}
func getfileSize(file string) int64 {
	//var filesize int
	var siZe int64
	f, err := os.Open(file)
	checkErr(err)
	defer f.Close()
	a, _ := f.Stat()
	size := a.Size()
	siZe = size / 1024
	num := size % 1024
	if num == 0 {
		fmt.Printf("文件大小：%d KB\n", siZe)
	} else {
		siZe += 1
		fmt.Printf("文件大小：%d KB\n", siZe)
	}

	return siZe
}
func byteString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 将文件名， 文件大小及文件的sha1值写入dirTreefile.txt文件。
func GetfileName(listfile []string) {
	//var filesize []byte
	sumFiles := len(listfile) // 文件总数
	file, err := os.Create(OutputFileName)
	checkErr(err)
	//file, err := os.OpenFile(OutputFileName, os.O_RDWR|os.O_CREATE, 0777)
	defer file.Close()
	fmt.Fprint(file, "本次共找到 ")
	fmt.Fprintf(file, "%d", sumFiles)
	fmt.Fprintln(file, " 个文件")
	//logger := log.New(file, "\r\n", log.Ldate)
	for i := 0; i < sumFiles; i++ {
		name := filepath.Base(listfile[i])
		sha1 := SHA1.GetSha1(listfile[i])
		fmt.Println(sha1)
		// 将sha1[20]byte 数组转化为 sha1[20]string
		//str_sha1 := byteString(sha1)
		filesize := getfileSize(listfile[i])
		//logger.Println(":" + name + ",")
		file.WriteString("Filename:" + name + ", Sha1:")
		fmt.Fprintf(file, "%v", sha1)
		file.WriteString(", fileSize:")
		//fmt.Fprintf(OutputFileName, "%d", filesize)
		file.WriteString(strconv.Itoa(int(filesize)) + "KB")
		file.WriteString(" \n")
	}
	fmt.Println("Done! 总共的文件数为：", sumFiles)
	listfile = make([]string, 0)
}
