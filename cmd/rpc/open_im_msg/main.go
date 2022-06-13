package main

import (
	"flag"
	rpcChat "open-im/internal/rpc/msg"
)

func main() {
	rpcPort := flag.Int("port", 10300, "rpc listening port")
	flag.Parse()
	rpcServer := rpcChat.NewRpcChatServer(*rpcPort)
	rpcServer.Run()
}
