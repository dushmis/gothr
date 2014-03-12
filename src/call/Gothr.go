package call

import (
	"fmt"
	"sync"
)

var (
	count int=0
	read sync.Mutex
	GO_COUNT_MAX int
	keeprunning bool=true
)

type Gothr struct{
}

func (this *Gothr) End(){
	fmt.Printf("%s\n","End Request")
	keeprunning=false
}

func (this *Gothr) Init(count_max int){
	GO_COUNT_MAX=count_max
}

func (this *Gothr) Execute(ev Curt){
	for {
			ch := make(chan int)
			go ac_st(ch,ev)
			<-ch
			if !keeprunning {
				break
			}
			fmt.Printf("%d\n",get_count())
	}
}


func DoStuff(c chan int,ev Curt){
	read.Lock()
	count+=1
	ev()
	c <- count;
	read.Unlock()
}


func get_count() int{
	return count
}


func ac_st(c chan int,ev Curt){
	ch := make(chan int,GO_COUNT_MAX)
	for ic := 1; ic <= GO_COUNT_MAX; ic ++ {
		fmt.Print("calling\n")
		go DoStuff(ch,ev)
	}
	<- ch
	c <- 1
}
