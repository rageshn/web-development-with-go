package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { //Infinite loop which listens to the connection
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		go handle(conn)
	}
}
