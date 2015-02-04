package main
import(
	"net"
	"fmt"
	"runtime"
	"os"
	"time"
)
const MY_IP= "129.241.187.150"
const MY_PORT = "20006"
const TARGET_PORT = "33546"
const TARGET_IP = "129.241.187.136"

func connectTo(ipAdr string, port string){

	serverAddr, err := net.ResolveTCPAddr("tcp",ipAdr+":"+port)
	checkError(err)

	con, err := net.DialTCP("tcp", nil, serverAddr);	
	checkError(err)

	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)	
	stop :=0
	
	msg:= "Connect to: "+MY_IP+":"+MY_PORT+"\x00"
	con.Write([]byte(msg))	
	
	for stop !=-1{
		input := ""
		fmt.Scanf("%s",&input)
		if input=="stop"{
			stop=-1
		}
		_,err := con.Write([]byte(input+"\x00"))
		checkError(err)


		_,err = con.Read(cd)
		checkError(err)

		fmt.Printf("%s\n",cd)
	}
}

func ListenerCon(port string){
	psock, err := net.Listen("tcp", ":"+port)
	checkError(err)
	
 	conn, err := psock.Accept()
 	checkError(err)
 	
 	go func(conn net.Conn){
		for {
			buf := make([]byte,1024)
			_,err := conn.Read(buf)
			checkError(err)
			fmt.Printf("%s\n",buf)
		}
	}(conn)
 	
  	for{  	
		msg:= "Boobies\x00"
		_,err := conn.Write([]byte(msg))
		checkError(err)
		time.Sleep(1*time.Second)
	 }		
}	

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go connectTo(TARGET_IP,TARGET_PORT)
	go ListenerCon(MY_PORT)	
	
	
	deadChan := make(chan int)
	<-deadChan
}
