package replays

import "github.com/jackc/pgx/v5/pgtype"

type CreateReplayParams struct {
	ReplayTitle string           `json:"replayTitle"`
	StageName   string           `json:"stageName"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
}
