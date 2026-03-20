package main

import (
	"fmt"
	presenceClient "gameapp/adapter/presence"
	"gameapp/adapter/redis"
	"gameapp/config"
	"gameapp/repository/redis/redismatching"
	"gameapp/scheduler"
	"gameapp/service/matchingservice"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	JwtSignKey = ""
)

func main() {
	// TODO - read config path from command line
	cfg := config.Load("config.yml")
	fmt.Printf("cfg: %+v\n", cfg)

	matchingSvc := setupServices(cfg)

	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		sch := scheduler.New(cfg.Scheduler, matchingSvc)

		wg.Add(1)
		sch.Start(done, &wg)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("received interrupt signal, shutting down gracefully..")
	done <- true
	time.Sleep(cfg.Application.GracefulShutdownTimeout)

	wg.Wait()
}

func setupServices(cfg config.Config) matchingservice.Service {
	redisAdapter := redis.New(cfg.Redis)

	matchingRepo := redismatching.New(redisAdapter)

	// TODO - add address to config
	presenceAdapter := presenceClient.New(":8086")

	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo, presenceAdapter, redisAdapter)

	return matchingSvc
}