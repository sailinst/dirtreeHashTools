package dirTree

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getdir(dirpath, suffix string) (files []string, err error) {
	files = make([]string, 0)
	//for i := 0; i < len(suffixArray); i++ {
	//	suffix := suffixArray[i]
	suffix = strings.ToUpper(suffix)
	//遍历目录
	err = filepath.Walk(dirpath, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error：", err)
			panic(err)
			return err
		}
		// 忽略目录
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}
func PrintFileDir(dirpath string) []string {

	fmt.Println("请输入不需要检索的文件类型和数量，即1,.go或者2,.go .txt ：")
	var sum int
	var suf string
	var suffix = make([]string, 0)
	var allfiles = make([]string, 0)
	fmt.Scanf("%d", &sum)
	for i := 0; i < sum; i++ {
		fmt.Scanln(&suf)
		suffix = append(suffix, suf)
	}
	sufList := delFileSuffix(suffix)
	//遍历目录中所有文件，将后缀名为suflist中有的，取出
	for _, v := range sufList {
		files, err := getdir(dirpath, v)
		if err != nil {
			fmt.Println("Error:", err)
		}
		for _, value := range files {
			allfiles = append(allfiles, value)
		}
	}
	fmt.Println("所有文件路径：", allfiles)
	return allfiles
}
func delFileSuffix(suffix []string) []string {
	var suffixArray = []string{"pptx", "pdfx", "xlsx", "docx", ".jpg", ".png", ".bmp", ".gif", ".jpeg", ".go", ".c", ".h", ".java", ".class", ".exe", ".js", "xml", "htm", ".html", ".css", ".txt", ".doc", ".xls", ".pdf", ".ppt", ".conf", ".log", ".ini", "sam", ".mp3", ".flash", ".avi", ".mp4", "wma", "rm"}
	suflist := list.New()
	var sufarray = make([]string, 0)
	// 将文件后缀名存入链表中，初始化链表
	for i := 0; i < len(suffixArray); i++ {
		suflist.PushBack(suffixArray[i])
	}
	for _, v := range suffix {
		for e := suflist.Front(); e != nil; e = e.Next() {
			if v == e.Value {
				suflist.Remove(e)
				fmt.Printf("从suffixArray[]中删除后缀 %s成功！ \n", v)
			} else if e.Next() == nil {
				fmt.Printf("suffixArray[]中不包含%s 的后缀名。\n", v)
			}
		}
	}
	for e := suflist.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(string)
		if !ok {
			fmt.Println("No string!")
			return nil
		}
		sufarray = append(sufarray, value)
	}
	return sufarray
}
