package main

import (
	"fmt"
	"net/rpc"
	"time"
)
const (
	serverAddress ="192.168.232.1:3339"//"114.212.83.131:3339"//
)
func main() {
	client, err := rpc.Dial("tcp", serverAddress )
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	a := "1234567"
	var reply time.Time
	var thetime time.Time
	old_time := time.Now()
	client.Call("Rpc_time.Timerpc",a,&reply)
	new_time := time.Now()

	d:=new_time.Sub(old_time)
	client.Close()
	//layout := "2006-01-02 3.04.05 PM"
	fmt.Println("系统原始时间",new_time)//.Format(layout))
	if reply!=thetime {
		fmt.Println("系统更新时间",reply.Add(d / 2))//.Format(layout))
	}else{
		fmt.Println("密码错误,无法同步更新")
	}
}


