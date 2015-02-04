package main

import (
	"net"
	"fmt"
	"runtime"
	"os"
)

const MY_IP= "129.241.187.150"
const TARGET_IP = "129.241.187.255"
const TARGET_PORT = "20006"
const LISTEN_PORT = "30000"

func ConnectTo(ipAdr string, port string){
	serverAddr, err := net.ResolveUDPAddr("udp",ipAdr+":"+port)
	checkError(err)

	con, err := net.DialUDP("udp", nil, serverAddr)
	checkError(err)
	
	stop := 0
		
	
	for stop !=-1{
		input := ""
		fmt.Scanf("%s",&input)
		if input=="stop"{
			stop=-1
		}
		con.Write([]byte(input+"\x00"))
		
	}
	fmt.Println("Closing UDP sender...")
}
func ListenerCon(port string){
	serverAddr, err := net.ResolveUDPAddr("udp",":"+port)
	checkError(err)

	psock, err := net.ListenUDP("udp4", serverAddr)
	checkError(err)

	buf := make([]byte,1024)
 
  	for {
		fmt.Println("Listening...") 
    		_, remoteAddr, err := psock.ReadFromUDP(buf)
		checkError(err)
    		if remoteAddr.IP.String() != MY_IP {
			fmt.Printf("%s\n",buf)
		}
	 }	
		
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())


	go ListenerCon(LISTEN_PORT)

	go ConnectTo(TARGET_IP,TARGET_PORT)
	go ListenerCon(TARGET_PORT)
	
	deadChan := make(chan int)
	<-deadChan
}

