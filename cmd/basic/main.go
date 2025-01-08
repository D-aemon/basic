package main

import (
	"basic/api/proto"
	"basic/internal/config"
	"basic/internal/db"
	"basic/internal/handler"
	"basic/internal/logger"
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"sync"
)

func startGrpcServer(wg *sync.WaitGroup, d db.DB, l *logger.Logger, grpcPort string) {
	defer wg.Done()
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	proto.RegisterStreamServiceServiceServer(s, &handler.StreamService{DB: d})
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", grpcPort))
	if err != nil {
		l.Error(fmt.Sprintf("failed to listen: %v", err))
		return
	}
	l.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%s", grpcPort))

	if err := s.Serve(lis); err != nil {
		l.Error(fmt.Sprintf("failed to serve: %v", err))
		return
	}
}

func startHttpServer(wg *sync.WaitGroup, l *logger.Logger, grpcPort, httpPort string) {
	defer wg.Done()
	// 2. 启动 HTTP 服务
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Error(fmt.Sprintf("Failed to dial server: %s", err.Error()))
		return
	}

	gwmux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher))
	// Register Greeter
	err = proto.RegisterStreamServiceServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		l.Error(fmt.Sprintf("Failed to register gateway: %s", err.Error()))
		return
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", httpPort),
		Handler: gwmux,
	}

	l.Info(fmt.Sprintf("Serving gRPC-Gateway on http://0.0.0.0:%s", httpPort))
	if err = gwServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		l.Error(fmt.Sprintf("HTTP server ListenAndServe: %s", err.Error()))
		return
	}
}

func main() {
	// 初始化 log
	l := logger.New(config.Cfg.Log.Level, config.Cfg.Log.Kinds, config.Cfg.Log.Project)
	// 链接数据库
	d, err := db.Connect(config.Cfg)
	if err != nil {
		l.Error(fmt.Sprintf("Failed to connect database, reason: %s", err.Error()))
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	// 启动 Grpc 服务
	go startGrpcServer(&wg, d, l, config.Cfg.Port.GrpcPort)
	// 启动 Http 服务
	go startHttpServer(&wg, l, config.Cfg.Port.GrpcPort, config.Cfg.Port.HttpPort)
	wg.Wait()
}
