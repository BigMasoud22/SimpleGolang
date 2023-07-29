package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

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
	//getting the ip address
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Something went wrong :%v", err.Error())
		connection.Close()
		os.Exit(1)
	}
	//if everything was alright we demonstrate the ip
	fmt.Printf("You are connected to the internet and your ip :%v", string(ip))
}
