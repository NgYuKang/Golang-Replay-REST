package replaycomments

type CreateReplayLikes struct {
	ReplayID       int    `json:"replayID" binding:"required"`
	CommentContent string `json:"commentContent" binding:"required"`
}
