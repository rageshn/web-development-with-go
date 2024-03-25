package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func foo(rw http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)

	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\n file: ", f, " Header: ", h, " ")

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `
	<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	</br>
	`+s)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
