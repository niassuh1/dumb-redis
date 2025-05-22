package main

import (
	"dumb-redis/pkg/serialization"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Server running on port :6379")

	l, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println(err)
	}

	conn, err := l.Accept()

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	for {

		resp := serialization.NewResp(conn)

		value, err := resp.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		fmt.Println(value.Array)

		conn.Write([]byte("+WORKS\r\n"))
	}
}
