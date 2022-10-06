package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/koind/anti-bruteforce/internal/bucket"
	"github.com/koind/anti-bruteforce/internal/config"
	"github.com/koind/anti-bruteforce/internal/list"
	"github.com/koind/anti-bruteforce/internal/server"
	"github.com/koind/anti-bruteforce/internal/service"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./configs/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cf := config.NewConfig(configPath)
	bs := bucket.NewStorage(cf)
	ls := list.NewStorage(cf)
	appService := service.NewService(bs, ls)
	grpcServer := server.NewServer(appService, cf)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			log.Printf("fail start server: %e", err)
		}
	}()

	<-ctx.Done()
	grpcServer.Stop()
}
