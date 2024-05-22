package replays

import "mime/multipart"

type CreateReplay struct {
	ReplayTitle string                `json:"replayTitle" form:"replayTitle" binding:"required"`
	StageName   string                `json:"stageName" form:"replayTitle" binding:"required"`
	ReplayFile  *multipart.FileHeader `json:"replayFile" form:"replayFile" binding:"required"`
}

type ListReplay struct {
	SortBy string `json:"sortBy" form:"sortBy"`
	Limit  int    `json:"limit" form:"limit"`
}
