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
	// Start client
	fmt.Println("Start client")
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection failled with %s\n", err)
	}

	for {
		//dec := gob.NewDecoder(conn)
		//dec.Decode()

		cmd, _ := bufio.NewReader(conn).ReadString('\n')
		cmd = strings.TrimSuffix(cmd, "\n")
		out, err := exec.Command(cmd).Output()
		
		if err != nil {
			fmt.Println("%s", err)
		}

		fmt.Println("Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)

		/*
		if runtime.GOOS == "windows" {
			cmd = exec.Command("dir")
		}
		*/

		/*
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		*/
		
		conn.Close()
		fmt.Println("Finished")
	}
}
