package main

import (
	"Golang-Replay-REST/api/replaycomments"
	"Golang-Replay-REST/api/replaylikes"
	"Golang-Replay-REST/api/replays"
	"Golang-Replay-REST/configs"
	"context"
	"flag"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dutchcoders/go-clamd"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	useFile := flag.Bool("useFile", true, "Whether to use env file")
	flag.Parse()
	var dbSource string
	var clamdURL string
	if *useFile {
		configs.LoadEnv()
		dbSource = configs.EnvDBSource()
		clamdURL = "tcp://127.0.0.1:3310"
	} else {
		dbSource = configs.EnvDBSourceDocker()
		clamdURL = configs.EnvClamDURL()
	}
	log.Println(dbSource)
	mode := configs.EnvGinMode()
	gin.SetMode(mode)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to connect to database on startup, exiting...")
	}

	defer conn.Close(context.Background())

	//S3
	awsSession := configs.ConnectAWS()
	uploadManager := s3manager.NewUploader(awsSession)
	downloadManager := s3manager.NewDownloader(awsSession)

	// clamd
	cav := clamd.NewClamd(clamdURL)

	router := gin.New()

	// Proxy from nginx
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	router.TrustedPlatform = "X-Forwarded-For"

	// gin logging and recovery stuff
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	rg := router.Group("/api")

	replays.SetupRoutes(rg, conn, uploadManager, downloadManager, cav)
	replaylikes.SetupRoutes(rg, conn)
	replaycomments.SetupRoutes(rg, conn)

	log.Fatal(router.Run(":8080"))

}
