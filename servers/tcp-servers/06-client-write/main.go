package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error while connecting: ", err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Hei.. I connected to you...")
}
