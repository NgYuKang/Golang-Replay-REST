package replays

import (
	"Golang-Replay-REST/api"
	"context"
	"fmt"

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
	"replayURL",
	"createdAt"
) VALUES (
	$1, $2, $3, $4
) RETURNING "replayID", "replayTitle", "stageName", "createdAt"
`

func (q *ReplayQueries) Create(ctx context.Context, arg CreateReplayParams) (Replay, error) {
	row := q.db.QueryRow(ctx, createReplay,
		arg.ReplayTitle,
		arg.StageName,
		arg.ReplayURL,
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

const listReplays = `--name: ListContacts
SELECT
    r."replayID",
    r."replayTitle",
    r."stageName",
    r."createdAt",
    COUNT(rl."likeID") as likes,
	COUNT(rc."commentID") as comments
FROM
    "replays" r
    LEFT JOIN "replayLikes" rl ON r."replayID" = rl."replayID"
	LEFT JOIN "replayComments" rc ON r."replayID" = rc."replayID"
GROUP BY
    r."replayID",
    r."replayTitle",
    r."stageName",
    r."createdAt"
ORDER BY
	%s DESC
LIMIT
	$1;
`

func (q *ReplayQueries) List(ctx context.Context, orderBy string, limit int) ([]Replay, error) {

	// Should not have sql injection: we manually set the orderBy string with a switch.
	// could still sanitize it...
	builtQuery := fmt.Sprintf(listReplays, orderBy)

	rows, err := q.db.Query(ctx, builtQuery, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	retList := []Replay{}

	for rows.Next() {
		var ret Replay
		if err := rows.Scan(
			&ret.ReplayID,
			&ret.ReplayTitle,
			&ret.StageName,
			&ret.CreatedAt,
			&ret.Likes,
			&ret.CommentCount,
		); err != nil {
			return nil, err
		}
		retList = append(retList, ret)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return retList, nil

}
