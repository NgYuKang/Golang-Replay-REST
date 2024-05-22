package replaylikes

import "github.com/jackc/pgx/v5/pgtype"

type ReplayLikes struct {
	ReplayID  int              `json:"replayID"`
	LikeID    int              `json:"likeID"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
}
