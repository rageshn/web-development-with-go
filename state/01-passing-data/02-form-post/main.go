package main

import (
	"io"
	"net/http"
)

func foo(rw http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	</br>
	`+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
