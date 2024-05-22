package replays

import (
	"log"
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

func (ctrl *ReplayController) CreateContact(ctx *gin.Context) {

	var payload *CreateContact

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
		log.Println(err)
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
