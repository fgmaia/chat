package services

import (
	"context"
	"fmt"
	"net"

	"github.com/fgmaia/chat/internal/controllers"
)

type udpServer struct {
	hostname       string
	port           string
	connectedUsers map[string]string
}

func NewUdpServer(hostname string,
	port string) Server {

	return &udpServer{hostname: hostname, port: port}
}

func (u *udpServer) Start(ctx context.Context) error {
	service := u.hostname + ":" + u.port

	udpAddr, err := net.ResolveUDPAddr("udp4", service)

	if err != nil {
		return err
	}

	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		return err
	}

	fmt.Println("UDP server up and listening on port " + u.port)

	defer ln.Close()

	actionHandler := controllers.NewControllerHandler()

	for {
		err := u.handleUDPConnection(ctx, ln, actionHandler)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (u *udpServer) handleUDPConnection(ctx context.Context, conn *net.UDPConn, actionHandler controllers.ControllerHandler) error {
	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if err != nil {
		return err
	}

	responseData, err := actionHandler.Handler(ctx, buffer[:n])
	if err != nil {
		return err
	}

	_, err = conn.WriteToUDP(responseData, addr)

	return err
}
