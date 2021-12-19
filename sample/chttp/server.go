package chttp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Request struct {

}

type Server struct {
	Addr string
	Port string
	Routes map[string]func(request Request)
}

func CreateServer(Addr string, Port string) *Server {
	return &Server{Addr: Addr, Port: Port}
}

func (server Server) Run(){
	listener, err := net.Listen("tcp", fmt.Sprint(server.Addr, ":", server.Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		method := strings.Fields(line)[0]
		route := strings.Fields(line)[1]
		fmt.Println(method)
		fmt.Println(route)
		break
	}
}