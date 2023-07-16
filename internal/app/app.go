package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	restHandler "github.com/begenov/tsarka-task/internal/delivery/http/v1"
	"github.com/begenov/tsarka-task/internal/repository/postgres"
	"github.com/begenov/tsarka-task/internal/repository/redis"
	"github.com/begenov/tsarka-task/internal/server"
	"github.com/begenov/tsarka-task/internal/service"
	"github.com/begenov/tsarka-task/pkg/database"

	redisClient "github.com/begenov/tsarka-task/pkg/redis"

	"github.com/begenov/tsarka-task/internal/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) error {
	// db
	db, err := database.Open(cfg.POSTGRES)
	if err != nil {
		return err
	}

	// redis client
	redisClient, err := redisClient.NewRedisClient(cfg.REDIS)
	if err != nil {
		return err
	}

	// Repository
	userRepo := postgres.NewUserRepo(db)
	counterRepo := redis.NewCouterRepo(redisClient)

	// Service
	userService := service.NewUserService(userRepo)
	counterService := service.NewCounterervice(counterRepo)

	// Handlers
	substrRestHandler := restHandler.NewSubstrHandler()
	emailAndInnRestHandler := restHandler.NewEmailInnHandler()
	counterRestHandler := restHandler.NewCounterHandler(counterService)
	userRestHandler := restHandler.NewUserHandler(userService)
	hashRestHandler := restHandler.NewHashHandler()

	// Routes
	router := gin.Default()
	api := router.Group("api/v1/rest")
	substrRestHandler.LoadRoutes(api)
	emailAndInnRestHandler.LoadRoutes(api)
	counterRestHandler.LoadRoutes(api)
	userRestHandler.LoadRoutes(api)
	hashRestHandler.LoadRoutes(api)

	// server
	srv := server.NewServer(cfg.HTTP, router)

	// run server
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("server run error: %v", err)
		}
	}()

	log.Printf("Server started %v", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	// stop server
	if err := srv.Stop(ctx); err != nil {
		log.Printf("failed to stop server: %v", err)
	}

	return nil
}
