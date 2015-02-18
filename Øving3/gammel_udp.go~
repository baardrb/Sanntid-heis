
package main

import (
	"net"
	"fmt"
	"runtime"
)

const(
	recieve_port = ":30000"
	send_port = ":20003"
	my_IP = "129.241.187.150"
	server_IP = "129.241.187.136"
)

func recieve(){

	buffer := make([]byte, 1024)


	addr, _ := net.ResolveUDPAddr("udp",send_port)
	sock, _ := net.ListenUDP("udp", addr)

	for{
		_, _,error := sock.ReadFromUDP(buffer)
		if error != nil{
			fmt.Println(error)
		}S
		fmt.Println(string(buffer[:]))
	}
}

func transmit(){

	addr, err := net.ResolveUDPAddr("udp", server_IP+send_port)
	handleError(err)

	sock, err := net.DialUDP("udp",nil,addr)
	handleError(err)

	message := ""
	fmt.Scanf("%s", &message)

	n, err := sock.Write([]byte(message+"\x00"))
	handleError(err)

	fmt.Println("På en skala fra 1 til teit er jeg: ", n)

}

func handleError(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	transmit()

	recieve()

}

