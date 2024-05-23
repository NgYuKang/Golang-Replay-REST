package replaycomments

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "not found",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "internal server err",
			})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":      http.StatusCreated,
		"message":     "created",
		"replayLikes": ret,
	})

}
