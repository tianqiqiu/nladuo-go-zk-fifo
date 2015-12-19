package main

import (
	"fmt"
	"github.com/nladuo/go-zk-fifo/fifo"
)

var (
	hosts    []string = []string{"127.0.0.1:2181"} // the zk server list
	basePath string   = "/fifo"                    // the application znode, you can create it by your self
	fifoData []byte   = []byte("the fifo data")    // the data of application's znode
	prefix   string   = "seq-"                     // the fifo prefix
)

func main() {
	// create zk connection
	err := fifo.EstablishZkConn(hosts)
	if err != nil {
		panic(err)

	}
	//create the distributed fifo
	myfifo := fifo.NewFifo(basePath, fifoData, prefix)
	//push data into fifo
	myfifo.Push("go-zk-fifo")
	//pop data from fifo
	data := myfifo.Pop()
	fmt.Println(data)
}