package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Received: %s", buf)
}

func main(){
	fmt.Println("GoTris")
	listen, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }
	for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}