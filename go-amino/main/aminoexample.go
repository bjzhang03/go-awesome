package main

import "fmt"
import "github.com/tendermint/go-amino"

func main() {
	//错误恢复机制
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered:", e)
		}
	}()
	// 消息定义
	type Message interface{}
	// 消息的封装
	type bcMessage struct {
		Message string
		Height  int
	}
	// 消息的响应
	type bcResponse struct {
		Status  int
		Message string
	}
	// 当前的status
	type bcStatus struct {
		Peers int
	}

	var cdc = amino.NewCodec()
	cdc.RegisterInterface((*Message)(nil), nil)
	cdc.RegisterConcrete(&bcMessage{}, "bcMessage", nil)
	cdc.RegisterConcrete(&bcResponse{}, "bcResponse", nil)
	cdc.RegisterConcrete(&bcStatus{}, "bcStatus", nil)
	// 消息的构造函数
	var bm = &bcMessage{Message: "ABC", Height: 100}
	var msg = bm

	var bz []byte // the marshalled bytes.
	var err error
	bz, err = cdc.MarshalBinary(msg)
	fmt.Printf("Encoded: %X (err: %v)\n", bz, err)

	var msg2 Message
	err = cdc.UnmarshalBinary(bz, &msg2)
	fmt.Printf("Decoded: %v (err: %v)\n", msg2, err)
	var bm2 = msg2.(*bcMessage)
	fmt.Printf("Decoded successfully: %v\n", *bm == *bm2)

}
