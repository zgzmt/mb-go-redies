package server_conn

import (
	"context"
	"fmt"
	logger "mb-go-redis/pkg/logs"
	"log"
	config "mb-go-redis/configs"
	server "mb-go-redis/app/interface/server"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//服务短连接处理包

func ListenAndServe(cfg *config.ServerConfig, handler server.Handler){
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil{
		log.Logger.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	//监听中断信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sigNot := <- signalChan
		switch sigNot {
		case syscall.SIGHUP, syscall.SIGQUIT:
			logger.G_Logger.Println("shuting down...")
			//关闭监听，阻止新连接进入
			listener.Close()
			//to_do 如何把剩余任务处理完之后再释放连接？
			handler.Closer()
		}
	}()

	logger.G_Logger.Printf("bind address: %s, listening...", cfg.Address)
	ctx, _ := context.WithCancel(context.Background())
	var waitG sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.G_Logger.Printf("accept err: ", err)
			continue
		}
		//这部分感觉很有问题，但是不知道问题在哪
		//waitG这样用好像没有意义
		go func() {
			defer func() {
				waitG.Done()
			}()
			waitG.Add(1)
			handler.Handle(ctx, conn)
		}()
	}
}
