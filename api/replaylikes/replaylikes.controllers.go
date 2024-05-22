package replaylikes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReplayLikesController struct {
	ReplayLikesDB *ReplayLikesQuery
}

func NewController(db *ReplayLikesQuery) *ReplayLikesController {
	return &ReplayLikesController{db}
}

func (ctrl *ReplayLikesController) Create(ctx *gin.Context) {

	var payload *CreateReplayLikes

	// CHANGE TO MULTIPART LATER
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed payload",
		})
		return
	}

	timeNow := time.Now()

	args := CreateReplayLikesParams{
		ReplayID: payload.ReplayID,
	}
	args.CreatedAt.Scan(timeNow)

	ret, err := ctrl.ReplayLikesDB.Create(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed insert",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":      http.StatusCreated,
		"message":     "created",
		"replayLikes": ret,
	})

}
