package main

import (
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

}
