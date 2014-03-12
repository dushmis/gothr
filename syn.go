package main

import (
	"call"
	"fmt"
	"time"
)

var (
	count int
)

func main() {
	c := &call.Gothr{}
	c.Init(20)
	c.Execute(Scan)

}

func Scan() int {
	count += 1
	fmt.Printf("%d\n", count)

	time.Sleep(1*time.Millisecond)
	return 9
}
