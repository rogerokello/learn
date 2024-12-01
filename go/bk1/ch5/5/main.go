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

	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

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

func getFile(name string) (*os.File, func(), error) {
	f, error := os.Open(name)

	if error != nil {
		return nil, nil, error
	}

	return f, func() {
		f.Close()
	}, nil
}
