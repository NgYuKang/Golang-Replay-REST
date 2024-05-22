package replays

type CreateReplay struct {
	ReplayTitle string `json:"replayTitle" binding:"required"`
	StageName   string `json:"stageName" binding:"required"`
	// ADD MULTIPART FILE FOR REPLAY LATER
}

type ListReplay struct {
	SortBy string `json:"sortBy" form:"sortBy"`
	Limit  int    `json:"limit" form:"limit"`
}
