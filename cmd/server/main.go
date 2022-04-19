package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fgmaia/chat/internal/infra/services"
)

func main() {
	chQuit := make(chan os.Signal, 2)
	signal.Notify(chQuit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for range chQuit {
			cancel()
			os.Exit(0)
		}
	}()

	start(ctx)
}

func start(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering from panic")
		}
	}()

	err := services.NewUdpServer("localhost", "8801").Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
