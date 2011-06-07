package main

import (
	"testing"
	"log"
	"http"
	"os"
)

type ResponseWriter struct {

}

func (rw *ResponseWriter) Header() http.Header {
	return nil
}
func (rw *ResponseWriter) Write(b []byte) (int, os.Error) {
	log.Println(b)
	return 0, nil
}
func (rw *ResponseWriter) WriteHeader(int) {
}

func TestIndexLen(t *testing.T) {
	res := new(ServerResource)
	resp := *new(ResponseWriter)
	log.Println(res)
	res.Index(resp)
}
