package main

import (
	"io"
	"os"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must supply atleast two arguments ")
	}

	file_name := os.Args[1]

	f, error := os.Open(file_name)
	if error != nil {
		log.Fatal(error)
	}
	defer f.Close()

	buffer := make([]byte, 2048)

	for {
		count, error := f.Read(buffer)
		os.Stdout.Write(buffer[:count])

		if error != nil {
			if error != io.EOF {
				log.Fatal(error, count)
			}
			break
		}
	}

}