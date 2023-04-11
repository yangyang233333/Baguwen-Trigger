package main

import "net/http"

type svr int

func (s *svr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	w.Write([]byte("你好！"))
}

func main() {
	var s svr
	http.ListenAndServe("localhost:8080", &s)
}
