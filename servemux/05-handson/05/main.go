package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(con net.Conn) {
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Fprintf(con, "you said: %s\n", ln)
	}
	defer con.Close()
}

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
		go handle(con)
		fmt.Println("Code got here")
		fmt.Fprintln(con, "Connection closed")
	}
}
