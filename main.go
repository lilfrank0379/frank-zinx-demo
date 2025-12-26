package main

import (
	"fmt"
	"frank-zinx-demo/impl"
)

/*
  @author lilfrank
  @date   2025/12/24 15:50
*/

func main() {
	fmt.Println("server start")

	server := impl.NewServer("[frank-zinx 001]")
	server.Serve()
}
