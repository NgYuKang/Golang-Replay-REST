package replays

type CreateReplay struct {
	ReplayTitle string `json:"replayTitle" binding:"required"`
	StageName   string `json:"stageName" binding:"required"`
	// ADD MULTIPART FILE FOR REPLAY LATER
}
