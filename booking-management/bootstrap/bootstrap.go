package bootstrap

import (
	"booking-management/http/router"
	"booking-management/http/server"
	"booking-management/utils/config"
	"booking-management/utils/container"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(ctx context.Context) {
	conf, err := config.Parse()
	if err != nil {
		panic(err)
	}

	ctr, err := container.Resolve(conf)
	if err != nil {
		panic(err)
	}

	r := router.Init(ctr)

	server := server.NewHTTPServer(conf, r)

	go server.ListenAndServe(ctx)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-c

	Destruct(ctx, ctr, server)

	os.Exit(0)
}

func Destruct(ctx context.Context, ctr container.Containers, server server.HTTPServer) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	go server.Shutdown(ctx)

	<-ctx.Done()

	log.Println("service shutdown gracefully")
}
