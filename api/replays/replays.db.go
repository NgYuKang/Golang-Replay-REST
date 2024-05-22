package replays

import (
	"Golang-Replay-REST/api"
	"context"

	"github.com/jackc/pgx/v5"
)

func NewReplayQuery(db api.DBTX) *ReplayQueries {
	return &ReplayQueries{db: db}
}

type ReplayQueries struct {
	db api.DBTX
}

func (q *ReplayQueries) WithTx(tx pgx.Tx) *ReplayQueries {
	return &ReplayQueries{
		db: tx,
	}
}

const createReplay = `--CreateReplay
INSERT INTO replays(
	"replayTitle",
	"stageName",
	"createdAt"
) VALUES (
	$1, $2, $3
) RETURNING "replayID", "replayTitle", "stageName", "createdAt"
`

func (q *ReplayQueries) Create(ctx context.Context, arg CreateReplayParams) (Replay, error) {
	row := q.db.QueryRow(ctx, createReplay,
		arg.ReplayTitle,
		arg.StageName,
		arg.CreatedAt,
	)
	var retData Replay
	err := row.Scan(
		&retData.ReplayID,
		&retData.ReplayTitle,
		&retData.StageName,
		&retData.CreatedAt,
	)
	return retData, err
}
