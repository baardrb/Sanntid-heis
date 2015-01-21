package main

import(
	. "fmt"
	"runtime"
	"time"
)
var sum int = 0

func adding(){
	for i := 0; i <= 1000000; i++ {
		sum++
		}
	}

func subtracting(){
	for j := 0; j <= 1000000; j++{
		sum--
		}
	}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go adding()
	go subtracting()
	
	time.Sleep(100*time.Millisecond)
	Println(sum)
	}
