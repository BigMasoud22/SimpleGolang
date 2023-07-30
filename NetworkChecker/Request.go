package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)
type logger struct{}
func main() {
	//we attempt to connect to the google.com
	connection, err := net.Dial("tcp", "google.com:http")
	if err != nil {
		fmt.Println("You are NOT connect to the internet")
		connection.Close()
		os.Exit(1)
	}
	//then we use an api due to using the features which it presents
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		fmt.Printf("oops , something went wrong error:%v", err.Error())
		connection.Close()
		os.Exit(1)
	}
	defer resp.Body.Close()
	//getting the ip address and demonstrate it 
	//first way :

	//_ , err := io.ReadAll(resp.Body)
	//fmt.Printf("You are connected to the internet and your ip :")

	//second way :

	//io.Copy(os.Stdout , resp.Body)

	//third way :
	log := logger{}
	io.Copy(log,resp.Body)
}

func (logger) Write(bs []byte)(length int,err error){
	    //in this way ,we customized out own write function so that we could be able to modify it whenever we want 
	    fmt.Print(string(bs))
        return len(bs) , nil
}

