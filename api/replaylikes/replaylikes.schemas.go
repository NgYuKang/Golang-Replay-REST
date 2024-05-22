package replaylikes

type CreateReplayLikes struct {
	ReplayID int `json:"replayID" binding:"required"`
}
