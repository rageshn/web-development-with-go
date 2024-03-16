package main

import (
	"bufio"
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
		scanner := bufio.NewScanner(con)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" || len(ln) == 0 {
				break
			}
			fmt.Fprintf(con, "you said: %s\n", ln)
		}
		fmt.Println("Code got here")
		io.WriteString(con, "Connection closed")
		con.Close()
	}
}
