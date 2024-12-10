package main

import "fmt"

type LogicProvider struct {

}

func (l LogicProvider) Process(data string) string {
	return fmt.Sprintf("Logic is to print string, `%s`", data)
}

type Logic interface {
	Process(string) string
}

type Client struct {
	L Logic
}

func (c Client) process() {
	fmt.Println(c.L.Process("some string"))
}


func main() {
	c := Client {
		L: LogicProvider{},
	}
	c.process()
}