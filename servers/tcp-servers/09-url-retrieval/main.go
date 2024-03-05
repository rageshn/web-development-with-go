package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0] //method
			u := strings.Fields(ln)[1] // url
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", u)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang='en'><head><meta charset="UTF-8">
    <title>Hello world</title></head><body><strong>Hello</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
	respond(conn)
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
