package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Gggdevv/go-web-template/config"
	"github.com/Gggdevv/go-web-template/internal/server"
)

func main() {
	cfg := config.NewConfig()

	srv := server.NewServer(cfg)
	go func() {
		if err := srv.Start(); err != nil {
			panic("Server start error ...")
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.ShutDown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
