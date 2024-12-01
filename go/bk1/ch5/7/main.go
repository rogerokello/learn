package main

import (
	"fmt"
	"io"
	"log"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must supply a file name")
	}

	len, err := fileLen(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File length is ", len, " bytes")
}

func fileLen(name string) (int, error) {
	f, err := os.Open(name)

	if err != nil {
		return 0, err
	}

	defer f.Close()

	buffer := make([]byte, 2048)
	numBytes := 0
	for {
		count, err := f.Read(buffer)

		numBytes = numBytes + count

		if err != nil && err == io.EOF {
			// it means we have reached the end of the file
			err = nil
			break
		}

		if err != nil {
			break 
		}
	}

	return numBytes, err
}