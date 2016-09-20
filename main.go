package main

import (
	"dirtreeHash/dirTree"
	"dirtreeHash/outFile"
	"fmt"
	"os"
	"time"
)

func dowork() {
	var path string
	var a string
	fmt.Println("输入目录的路径：")
	fmt.Scan(&path)
	// 判断path是否有效， 若无效，则提醒用户重新输入path

	fileList := dirTree.PrintFileDir(path)
	//fileList := dirTree.PrintFileList(path)
	outFile.GetfileName(fileList)

	//fmt.Println("文件路径为：", fileList)
	fmt.Printf("输入quit退出，输入其他则继续：\t")
	fmt.Scan(&a)
	//fmt.scanln(&a)
	switch a {
	case "quit":
		exit()
	case "退出":
		exit()
	default:

		dowork()
	}

	// 初始化变量
	//a = ""
	fileList = make([]string, 0)
}
func exit() {
	fmt.Println("系统即将退出，谢谢使用！")
	time.Sleep(1000)
	os.Exit(1)
}

func main() {
	dowork()
}
