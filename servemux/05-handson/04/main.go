package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		io.WriteString(con, "Hello\n")
		con.Close()
	}
}
