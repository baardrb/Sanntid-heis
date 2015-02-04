package main

import (
	"net"
	"fmt"
	"runtime"
)

const(
	recieve_port = ":20003"
	send_port = ":33546"
	my_IP = "129.241.187.150"
	server_IP = "129.241.187.136"
)

func connect(){
	addr, err := net.ResolveTCPAddr("tcp", server_IP + send_port)
	checkError(err)

	laddr, _ := net.ResolveTCPAddr("tcp", my_IP + send_port)

	con, err := net.DialTCP("tcp",nil, addr)
	checkError(err)

	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s\n",cd)

	message := "Connect to: " + my_IP + recieve_port + "\x00"
	con.Write([]byte(message))

	input := ""
	fmt.Scanf("%s", &input)

	_,err = con.Write([]byte(input))
	checkError(err)

	_,err = con.Read(cd)
	checkError(err)

	fmt.Printf("%s\n",cd)  

}

func checkError(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	connect()

}