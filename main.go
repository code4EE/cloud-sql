package main

import (
	"context"
	"os"
	"os/signal"

	inits "github.com/code4EE/cloud-sql/inits"
	"github.com/code4EE/cloud-sql/server"
)

func init() {
	inits.InitAll()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
	server.RunServer(ctx)
}
