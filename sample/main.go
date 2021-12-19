package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main(){
	listener, err := net.Listen("tcp", "192.168.0.40:8080")
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