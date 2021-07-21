package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	//获取命令行参数
	if len(os.Args) < 2 {
		fmt.Println("参数有误！")
		fmt.Println("请输入如下格式:./ssh_shell.exe \"ls -al /var/root\"")
		return
	}
	command := os.Args[1]
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", "127.0.0.1:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("alpine")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println("SSH 连接失败!")
		return
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("创建iPhone会话失败!")
		return
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		fmt.Println("命令没有执行成功!")
		return
	}
	fmt.Println(b.String())
}
