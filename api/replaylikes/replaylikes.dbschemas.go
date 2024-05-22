package replaylikes

import "github.com/jackc/pgx/v5/pgtype"

type CreateReplayLikesParams struct {
	ReplayID  int              `json:"replayID"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
}
