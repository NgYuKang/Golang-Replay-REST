package replays

import "github.com/jackc/pgx/v5/pgtype"

type CreateReplayParams struct {
	ReplayTitle    string           `json:"replayTitle"`
	StageName      string           `json:"stageName"`
	ReplayFileName string           `json:"ReplayFileName"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
}
