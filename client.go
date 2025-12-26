package main

import (
	"fmt"
	"net"
	"time"
)

/*
  @author lilfrank
  @date   2025/12/24 17:09
*/

func main() {
	fmt.Println("client start")

	conn, err := net.Dial("tcp", "127.0.0.1:8686")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for {
		data := []byte("hello world")
		fmt.Println("client write:", string(data))

		for len(data) > 0 {
			n, err := conn.Write(data)
			if err != nil {
				fmt.Println("error:", err)
			}
			data = data[n:]
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println("client read:", string(buf[:cnt])+"\n")

		time.Sleep(6 * time.Second)
	}
}
