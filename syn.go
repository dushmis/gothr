package main


import (
	"fmt"
	"call"
	"time"
	
)


func main(){
	c:=&call.Gothr{}
	c.Init(20)
	go c.Execute(Scan)
	time.Sleep(10*time.Second)
	c.End()
}

func Scan() int{
	fmt.Printf("%s\n","------")
	return 9
}