package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServerHttp(http.ResponseWriter, *http.Request)

func main() {

}
