package replays

import (
	"Golang-Replay-REST/api/replaycomments"

	"github.com/jackc/pgx/v5/pgtype"
)

type Replay struct {
	ReplayID    int                             `json:"replayID"`
	ReplayTitle string                          `json:"replayTitle"`
	StageName   string                          `json:"stageName"`
	CreatedAt   pgtype.Timestamp                `json:"createdAt"`
	Likes       int                             `json:"likes"`
	Comments    []replaycomments.ReplayComments `json:"comments"`
}
