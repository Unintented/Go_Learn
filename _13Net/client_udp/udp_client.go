package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败,error:", err)
		return
	}
	defer udpConn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for true {
		var data [1024]byte
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := udpConn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("write to udp failed, error:", err)
			continue
		}
		n, addr, err := udpConn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed, error:", err)
			continue
		}
		fmt.Println("from:", addr, "data:", string(data[:n]))
	}

}
