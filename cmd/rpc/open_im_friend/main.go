package main

import (
	"flag"
	"open-im/internal/rpc/friend"
)

func main() {
	rpcPort := flag.Int("port", 10200, "get RpcFriendPort from cmd,default 12000 as port")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
