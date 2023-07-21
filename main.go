package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取网络接口失败:", err)
	}
	return

	//遍历网络接口信息，找到第一个非回环的ipv4地址
	for _, addr := range interfaces {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("主机的第一个非回环地址:", ipnet.IP.String())
				return
			} else if ipnet.IP.To16() != nil {
				fmt.Println("主机的第一个非回环的 IPv6 地址:", ipnet.IP.String())
				return
			}
		}
	}
	fmt.Println("未找到非回环的IP地址")
}

