package main

import (
	pb "auth/genproto/user"
	"auth/internal/api"
	"auth/internal/config"
	"auth/internal/service"
	"auth/internal/storage"
	"auth/internal/storage/postgres"
	"auth/internal/storage/redis"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to postgres: %v", err)
	}
	defer db.Close()

	rdb := redis.NewRedisDB(cfg)
	defer rdb.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go RunService(&wg, db, cfg.AUTH_SERVICE_PORT)
	go RunRouter(&wg, db, rdb, cfg)

	wg.Wait()
}

func RunService(wg *sync.WaitGroup, db storage.IStorage, port string) {
    defer wg.Done()

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("error while listening: %v", err)
    }
    defer lis.Close()

    server := grpc.NewServer(
        grpc.MaxRecvMsgSize(1024*1024), // 1 MB
        grpc.MaxSendMsgSize(1024*1024), // 1 MB
    )
    pb.RegisterUserServer(server, service.NewUserService(db))

    log.Printf("Service is listening on port %s...\n", port)
    if err := server.Serve(lis); err != nil {
        log.Fatalf("error while serving auth service: %s", err)
    }
}


func RunRouter(wg *sync.WaitGroup, db storage.IStorage, rdb *redis.RedisDB, cfg *config.Config) {
	defer wg.Done()

	router := api.NewRouter(db, rdb, cfg)

	log.Printf("Router is running on port %s...\n", cfg.AUTH_ROUTER_PORT)
	if err := router.Run(cfg.AUTH_ROUTER_PORT); err != nil {
		log.Fatalf("error while running router: %s", err)
	}
}
