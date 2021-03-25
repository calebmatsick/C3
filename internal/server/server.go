package main

import "net"
import "fmt"
import "bufio"

func main() {
    fmt.Println("Start server...")

    // Listen on port 8000
    ln, _ := net.Listen("tcp", ":8000")

    // Accept connection
    conn, _ := ln.Accept()

    // Loop until interrupt
    for {
        // Get message, output
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Message Recieved:", string(message))
    }
}