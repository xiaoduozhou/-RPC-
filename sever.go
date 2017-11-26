package main

import (
    "fmt"
    "net"
    "net/rpc"
    "os"
    "time"
)
type Rpc_time int

func (t *Rpc_time)Timerpc(a string, reply *time.Time) error{
	if(a == "123456"){
		*reply  = time.Now()
		fmt.Println(time.Now())
		return nil
	}
	reply = nil
    return  nil
}
func main() {
    rpc.Register(new(Rpc_time))
    rpc.HandleHTTP()

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":3339")
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go rpc.ServeConn(conn)
    }

}
func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}