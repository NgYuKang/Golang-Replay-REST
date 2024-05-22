package replaycomments

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReplayCommentsController struct {
	ReplayLikesDB *ReplayCommentsQuery
}

func NewController(db *ReplayCommentsQuery) *ReplayCommentsController {
	return &ReplayCommentsController{db}
}

func (ctrl *ReplayCommentsController) Create(ctx *gin.Context) {

	var payload *CreateReplayLikes

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed payload",
		})
		return
	}

	timeNow := time.Now()

	args := CreateReplayCommentsParams{
		ReplayID:       payload.ReplayID,
		CommentContent: payload.CommentContent,
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
