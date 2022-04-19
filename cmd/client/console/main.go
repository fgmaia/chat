package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/fgmaia/chat/internal/entities"
)

func main() {
	hostName := "localhost"
	portNum := "8801"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	payload := entities.Payload{
		Username:    "testeuser",
		Action:      "send",
		SendMessage: "teste blablla",
	}
	data, err := json.Marshal(&payload)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(data)

	if err != nil {
		log.Println(err)
	}

	// receive message from server
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP Server : ", addr)
	fmt.Println("Received from UDP server : ", string(buffer[:n]))

}
