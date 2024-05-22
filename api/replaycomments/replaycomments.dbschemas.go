package replaycomments

import "github.com/jackc/pgx/v5/pgtype"

type CreateReplayCommentsParams struct {
	ReplayID       int              `json:"replayID"`
	CommentContent string           `json:"commentContent"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
}
