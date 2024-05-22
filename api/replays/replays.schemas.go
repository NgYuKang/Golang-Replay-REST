package replays

type CreateContact struct {
	ReplayTitle string `json:"replayTitle" binding:"required"`
	StageName   string `json:"stageName" binding:"required"`
	// ADD MULTIPART FILE FOR REPLAY LATER
}
