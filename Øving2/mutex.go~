package main

import(
	. "fmt"
	"runtime"
	"time"
)
var sum int = 0

var buffer = make(chan int,1)
var done_add = make(chan int,1)
var done_sub = make(chan int,1)


func adding(){
	for i := 0; i <= 1000000; i++ {
		<- buffer 	
		sum++
		buffer <- 1
		}
	done_add <- 1
	}

func subtracting(){
	for j := 0; j <= 1000001; j++{
		<- buffer 	
		sum--
		buffer <- 1
		}
	done_sub <- 1
	}

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU())
	buffer <- 1
	go adding()
	go subtracting()
	
	time.Sleep(100*time.Millisecond)
	<-done_add	
	<-done_sub
	Println(sum)
	}
