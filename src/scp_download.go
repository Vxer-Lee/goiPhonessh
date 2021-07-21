package main

import (
	"fmt"
	"os"
	"path/filepath"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func download_remote_file(remotefile string) {

	_, fileName := filepath.Split(remotefile)
	//用密码连接上iPhone的SSH
	clientConfig, _ := auth.PasswordKey("root", "alpine", ssh.InsecureIgnoreHostKey())
	//创建一个scp客户端
	client := scp.NewClient("127.0.0.1:22", &clientConfig)
	//与iPhone建立连接
	err := client.Connect()
	if err != nil {
		panic("连接SSH失败,请检查端口是否映射或者密码正确!")
		return
	}
	fmt.Println("与iPhone SSH连接成功!")
	//准备下载文件
	fmt.Println("准备下载文件!")
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("请以管理员权限运行此程序!")
		return
	}
	defer f.Close()

	err = client.CopyFromRemote(f, remotefile)
	if err != nil {
		fmt.Printf("下载%s文件失败,请自行ssh到iPhone查看路径是否有误!\n", remotefile)
		return
	}
	fmt.Printf("%s文件下载成功!\n", fileName)
}

func main() {
	//获取命令行参数
	if len(os.Args) != 2 {
		fmt.Println("参数有误！")
		fmt.Println("请输入如下格式:./scp_download.exe /var/root/keychain.txt")
		return
	}

	remotefile := os.Args[1]
	download_remote_file(remotefile)
}
