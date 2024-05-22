package replays

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReplayController struct {
	ReplayDB *ReplayQueries
}

func NewController(replayDB *ReplayQueries) *ReplayController {
	return &ReplayController{replayDB}
}

func (ctrl *ReplayController) Create(ctx *gin.Context) {

	var payload *CreateReplay

	// CHANGE TO MULTIPART LATER
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed payload",
		})
		return
	}

	// Virus Scan the file

	// Wait for completion

	// Check Results, return err if malicious

	// Encrypt file and upload to aws s3

	// Get link

	timeNow := time.Now()

	args := CreateReplayParams{
		ReplayTitle: payload.ReplayTitle,
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
