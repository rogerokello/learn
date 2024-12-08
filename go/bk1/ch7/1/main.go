package main

import (
	"fmt"
	"time"
)



type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total = c.total + 1
	c.lastUpdated = time.Now()
}

func (c *Counter) String() string{
	return fmt.Sprintf("Total %d was last updated at %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	doWrongUpdate(c)
	fmt.Println("Wrong update", c.String())
	var cPtr *Counter = &c
	doRightUpdate(cPtr)
	fmt.Println("Right update", c.String())
	// cPtr.Increment()
	// fmt.Println(cPtr.String())

}

func doWrongUpdate(c Counter) {
	c.Increment()
}

func doRightUpdate(c *Counter) {
	c.Increment()
}
