package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	log.Println("Starting http server...")

	err := os.Setenv("VERSION", "4.0")
	if err != nil {
		panic("set env panic")
	}

	s := NewServer()

	s.registerHandler("/healthz", handleHealthz)
	s.registerHandler("/time", handleTime)

	err = s.Start(":80")
	if err != nil {
		panic("start http server error")
	}

}

// /healthz handler
func handleHealthz(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)

	log.Println("hint /healthz, from " + getIp(req) + ", statusCode is 200")
}

// /time handler
func handleTime(resp http.ResponseWriter, req *http.Request) {
	statusCode := 200
	resp.WriteHeader(statusCode)
	_, err := resp.Write([]byte(time.Now().String()))
	if err != nil {
		statusCode = 500
		resp.WriteHeader(statusCode)
	}

	log.Println("hint /time, from " + getIp(req) + ", statusCode is " + strconv.Itoa(statusCode))
}

// get ip address
func getIp(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		return forwarded
	}

	return req.RemoteAddr
}
