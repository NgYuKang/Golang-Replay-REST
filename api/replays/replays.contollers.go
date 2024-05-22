package replays

import (
	"Golang-Replay-REST/configs"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type ReplayController struct {
	ReplayDB   *ReplayQueries
	uploader   *s3manager.Uploader   // Can be interfaced to allow testing
	downloader *s3manager.Downloader // can be interfaced to allow testing
}

func NewController(replayDB *ReplayQueries, uploader *s3manager.Uploader, downloader *s3manager.Downloader) *ReplayController {
	return &ReplayController{replayDB, uploader, downloader}
}

func (ctrl *ReplayController) Create(ctx *gin.Context) {

	var payload *CreateReplay

	// CHANGE TO MULTIPART LATER
	if err := ctx.ShouldBindWith(&payload, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed payload",
		})
		return
	}

	replayFile, err := payload.ReplayFile.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed open replay",
		})
	}

	// Virus Scan the file

	// Wait for completion

	// Check Results, return err if malicious

	// Encrypt file

	// Upload and get link
	timeNow := time.Now()
	filename := fmt.Sprintf("%s-%s", timeNow.Format("2006-01-02"), uuid.New())
	resUpload, err := ctrl.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(configs.EnvAWSBucket()),
		Key:    aws.String(filename),
		Body:   replayFile,
	})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed upload replay",
		})
		return
	}

	uploadedURL := resUpload.Location

	args := CreateReplayParams{
		ReplayTitle: payload.ReplayTitle,
		ReplayURL:   uploadedURL,
		StageName:   payload.StageName,
	}
	args.CreatedAt.Scan(timeNow)

	replay, err := ctrl.ReplayDB.Create(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed insert",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "created",
		"replay":  replay,
	})

}

func (ctrl *ReplayController) List(ctx *gin.Context) {

	var payload ListReplay

	// CHANGE TO MULTIPART LATER
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed payload",
		})
		return
	}

	var sortByStr string
	switch payload.SortBy {
	case "createdAt":
		sortByStr = "r.\"createdAt\""
	case "likes":
		sortByStr = "likes"
	default:
		sortByStr = "r.\"createdAt\""
	}

	if payload.Limit == 0 {
		payload.Limit = 10
	}

	replays, err := ctrl.ReplayDB.List(ctx, sortByStr, payload.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed to retrieve list of replays",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully retrieved list of replays",
		"replays": replays,
	})
}
