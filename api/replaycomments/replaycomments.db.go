package replaycomments

import (
	"Golang-Replay-REST/api"
	"context"

	"github.com/jackc/pgx/v5"
)

func NewReplayCommentsQuery(db api.DBTX) *ReplayCommentsQuery {
	return &ReplayCommentsQuery{db: db}
}

type ReplayCommentsQuery struct {
	db api.DBTX
}

func (q *ReplayCommentsQuery) WithTx(tx pgx.Tx) *ReplayCommentsQuery {
	return &ReplayCommentsQuery{
		db: tx,
	}
}

const createReplayLikes = `--CreateReplayComments
INSERT INTO "replayComments"(
	"replayID",
	"commentContent",
	"createdAt"
) VALUES (
	$1, $2, $3
) RETURNING "replayID", "commentID", "commentContent", "createdAt"
`

func (q *ReplayCommentsQuery) Create(ctx context.Context, arg CreateReplayCommentsParams) (ReplayComments, error) {
	row := q.db.QueryRow(ctx, createReplayLikes,
		arg.ReplayID,
		arg.CommentContent,
		arg.CreatedAt,
	)
	var retData ReplayComments
	err := row.Scan(
		&retData.ReplayID,
		&retData.CommentID,
		&retData.CommentContent,
		&retData.CreatedAt,
	)
	return retData, err
}
