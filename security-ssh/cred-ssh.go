package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var username = "ritesh.chauhan"
var password = "S"
var host = "<hostname>"
var command = "hostname"

func main(){
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil{
		log.Fatal("Error dialing to the server : ", err)
	}
	log.Println(string(client.ClientVersion()))

	session, err := client.NewSession()
	if err != nil{
		log.Fatal("Error starting the ssh session: ", err)
	}
	defer session.Close()
	session.Stdout = os.Stdout

	err = session.Run(command)
	if err != nil{
		log.Fatal("Error Running command : ", err)
	}
}