package transfer

import (
	// Standard
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	// C3
	"github.com/calebmatsick/C3/pkg/security"
)


const BUFFERSIZE = 1024


func fillString(returnString string, toLength int) string {
	for {
		lengthString := len(returnString)
		if lengthString < toLength {
				returnString = returnString + ":"
				continue
		}
		break
	}
	return returnString
}


func sendFile(conn net.Conn, givenFile string) {
	fmt.Println("")
	enc := gob.NewEncoder(conn)
	file, err := os.Open(givenFile)

	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	encFileSize := security.Encrypt(fileSize)
	enc.Encode(encFileSize)
	encFileName := security.Encrypt(fileName)
	enc.Encode(encFileName)
	sendBuffer := make([]byte, BUFFERSIZE)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		encSendBuffer := security.Encrypt(string(sendBuffer))
		enc.Encode(encSendBuffer)
	}
	fmt.Println("File has been sent")
}


func recieveFile(conn net.Conn) {
	dec := gob.NewDecoder(conn)

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	fileName:= strings.Trim(string(bufferFileName), ":")


	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	var recievedBytes int64

	for {
		if (fileSize - recievedBytes) < BUFFERSIZE {
			io.CopyN(newFile, conn, (fileSize - recievedBytes))
			dec.Decode(make([]byte, (recievedBytes+BUFFERSIZE)))
			break
		}
		io.CopyN(newFile, conn, BUFFERSIZE)
		recievedBytes += BUFFERSIZE
	}
}