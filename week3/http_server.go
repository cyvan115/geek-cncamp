package main

import (
	"log"
	"net/http"
	"os"
)

type MyHttp struct {
	interceptors   []MyInterceptor
	handlerMapping map[string]http.HandlerFunc
}

// MyInterceptor fake interceptor, only pre-process
type MyInterceptor func(http.ResponseWriter, *http.Request) bool

// NewServer make a new http server struct
func NewServer() *MyHttp {
	return &MyHttp{
		interceptors:   []MyInterceptor{httpHeaderInterceptor, osEnvInterceptor},
		handlerMapping: make(map[string]http.HandlerFunc),
	}
}

// Start starts the server
func (server *MyHttp) Start(addr string) error {
	for pattern, handler := range server.handlerMapping {
		handler := handler

		// fake enhance
		enhanced := func(resp http.ResponseWriter, req *http.Request) {
			for _, interceptor := range server.interceptors {
				ok := interceptor(resp, req)
				if ok != true {
					_, err := resp.Write([]byte("internal error"))
					if err != nil {
						resp.WriteHeader(500)
						log.Println("hint " + pattern + ", from " + getIp(req) + ", statusCode is 500")
						return
					}
				}
			}

			handler(resp, req)
		}

		// register handler
		http.HandleFunc(pattern, enhanced)
	}

	// listen and serve
	return http.ListenAndServe(addr, nil)
}

// register interceptor
func (server *MyHttp) registerInterceptor(foo MyInterceptor) {
	server.interceptors = append(server.interceptors, foo)
}

// register http handler
func (server *MyHttp) registerHandler(pattern string, handler http.HandlerFunc) {
	server.handlerMapping[pattern] = handler
}

// copy header from request to response
func httpHeaderInterceptor(resp http.ResponseWriter, req *http.Request) bool {
	for headerKey, headerValue := range req.Header {
		for _, item := range headerValue {
			resp.Header().Add(headerKey, item)
		}
	}

	return true
}

// get env variable VERSION and write to resp header
func osEnvInterceptor(resp http.ResponseWriter, req *http.Request) bool {
	version := os.Getenv("VERSION")
	resp.Header().Add("OS_ENV_VERSION", version)

	return true
}
