package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func foo(rw http.ResponseWriter, req *http.Request) {
	ck, err := req.Cookie("session-id")
	if err != nil {
		fmt.Println("Error: ", err)
		id := uuid.NewV4()
		ck = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(rw, ck)
	}
	fmt.Println(ck)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
