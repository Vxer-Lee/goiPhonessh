package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func run_shell(remotepath string) {
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", "127.0.0.1:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("alpine")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return
	}
	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		return
	}
	defer session.Close()
	cmd := "ls -al "
	cmd += remotepath
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return
	}
	fmt.Println(b.String())
}
func upload_file2remote(localfile string, remotefile string) {
	_, fileName := filepath.Split(localfile)
	remotepath, _ := filepath.Split(remotefile)
	//用密码连接上iPhone的SSH
	clientConfig, _ := auth.PasswordKey("root", "alpine", ssh.InsecureIgnoreHostKey())
	//创建一个scp客户端
	client := scp.NewClient("127.0.0.1:22", &clientConfig)
	//与iPhone建立连接
	err := client.Connect()
	if err != nil {
		fmt.Println("连接SSH失败,请检查端口是否映射或者密码正确!")
		return
	}
	fmt.Println("与iPhone SSH连接成功!")
	f, err := os.OpenFile(localfile, os.O_RDWR, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("本地文件被占用!")
		return
	}
	defer f.Close()

	//准备上传文件
	fmt.Println("准备上传文件")
	//上传程序到iPhone
	err = client.CopyFile(f, remotefile, "0655")
	if err != nil {
		fmt.Println("抱歉！上传文件失败了。请检查路径是否有误？")
		return
	}
	fmt.Printf("文件%s成功上传到iPhone->[%s]!\n\n", fileName, remotefile)
	run_shell(remotepath)
}

func main() {
	//获取命令行参数
	if len(os.Args) != 3 {
		fmt.Println("参数有误！")
		fmt.Println("请输入如下格式:./scp_upload.exe ./keychain.txt /var/root/keychain.txt")
		return
	}
	localfile := os.Args[1]
	remotefile := os.Args[2]
	upload_file2remote(localfile, remotefile)
}
