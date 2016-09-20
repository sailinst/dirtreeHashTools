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
	fmt.Scanln(&path)
	// 判断path是否有效， 若无效，则提醒用户重新输入path

	fileList := dirTree.PrintFileDir(path)
	//fileList := dirTree.PrintFileList(path)
	outFile.GetfileName(fileList)

	//fmt.Println("文件路径为：", fileList)
	fmt.Printf("输入命令继续操作：\n \t\tgo:\t继续...\n \t\tquit：\t退出系统\n \t\t退出：\t退出系统\n")
LB:
	fmt.Scanln(&a)
	//fmt.scanln(&a)
	switch a {
	default:

		fmt.Println("输入有误，请重新输入...")
		goto LB
	case "go":
		dowork()
	case "quit":
		exit()
	case "退出":
		exit()

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
