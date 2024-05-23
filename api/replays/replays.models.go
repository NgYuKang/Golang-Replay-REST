package replays

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Struct for List only
type Replay struct {
	ReplayID     int              `json:"replayID"`
	ReplayTitle  string           `json:"replayTitle"`
	StageName    string           `json:"stageName"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	Likes        int              `json:"likes"`
	CommentCount int              `json:"commentCount"`
	// Comments    []replaycomments.ReplayComments `json:"comments"` // NOT NEEDED FOR LIST
}
