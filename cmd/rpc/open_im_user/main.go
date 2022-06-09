package main

import (
	"flag"
	"open-im/internal/rpc/user"
)

func main() {
	rpcPort := flag.Int("port", 10100, "rpc listening port")
	flag.Parse()
	rpcServer := user.NewUserServer(*rpcPort)
	rpcServer.Run()
}
