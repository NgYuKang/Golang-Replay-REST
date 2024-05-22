package replaylikes

import (
	"Golang-Replay-REST/api"
	"context"

	"github.com/jackc/pgx/v5"
)

func NewReplayLikesQuery(db api.DBTX) *ReplayLikesQuery {
	return &ReplayLikesQuery{db: db}
}

type ReplayLikesQuery struct {
	db api.DBTX
}

func (q *ReplayLikesQuery) WithTx(tx pgx.Tx) *ReplayLikesQuery {
	return &ReplayLikesQuery{
		db: tx,
	}
}

const createReplayLikes = `--CreateReplayLikes
INSERT INTO "replayLikes"(
	"replayID",
	"createdAt"
) VALUES (
	$1, $2
) RETURNING "replayID", "likeID", "createdAt"
`

func (q *ReplayLikesQuery) Create(ctx context.Context, arg CreateReplayLikesParams) (ReplayLikes, error) {
	row := q.db.QueryRow(ctx, createReplayLikes,
		arg.ReplayID,
		arg.CreatedAt,
	)
	var retData ReplayLikes
	err := row.Scan(
		&retData.ReplayID,
		&retData.LikeID,
		&retData.CreatedAt,
	)
	return retData, err
}
