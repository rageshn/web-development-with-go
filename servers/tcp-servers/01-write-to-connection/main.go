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
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Error while accepting request: ", err)
		}
		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I Hope!")
		conn.Close()
	}
}
