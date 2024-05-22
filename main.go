package main

import (
	"Golang-Replay-REST/api/replaycomments"
	"Golang-Replay-REST/api/replaylikes"
	"Golang-Replay-REST/api/replays"
	"Golang-Replay-REST/configs"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	configs.LoadEnv()
	mode := configs.EnvGinMode()
	gin.SetMode(mode)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	conn, err := pgx.Connect(ctx, configs.EnvDBSource())
	if err != nil {
		log.Fatal("Failed to connect to database on startup, exiting...")
	}

	defer conn.Close(context.Background())

	router := gin.New()

	// Proxy from nginx
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	router.TrustedPlatform = "X-Forwarded-For"

	// gin logging and recovery stuff
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	rg := router.Group("/api")

	replays.SetupRoutes(rg, conn)
	replaylikes.SetupRoutes(rg, conn)
	replaycomments.SetupRoutes(rg, conn)

	log.Fatal(router.Run(":8080"))

}
