package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Error while setting deadline: ", err)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said: %s\n", ln)
	}
	defer conn.Close()

	print("************CODE ENDED***************")
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error while listening: ", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Error while accepting: ", err)
		}

		go handle(conn)
	}
}
