package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
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

	start := time.Now()
	for {
		count, error := f.Read(buffer)
		os.Stdout.Write(buffer[:count])

		if error != nil && error == io.EOF {
			break
		}

		if error != nil {
			log.Fatal(error, count)
		}
	}
	fmt.Println("it took ", time.Since(start).Seconds(), " seconds")

}