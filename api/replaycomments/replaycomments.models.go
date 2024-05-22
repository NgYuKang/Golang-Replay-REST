package replaycomments

import "github.com/jackc/pgx/v5/pgtype"

type ReplayComments struct {
	ReplayID       int              `json:"replayID"`
	CommentID      int              `json:"commentID"`
	CommentContent string           `json:"commentContent"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
}
