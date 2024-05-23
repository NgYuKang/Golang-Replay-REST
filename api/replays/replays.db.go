package replays

import (
	"Golang-Replay-REST/api"
	"Golang-Replay-REST/api/replaycomments"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	"replayFileName",
	"createdAt"
) VALUES (
	$1, $2, $3, $4
) RETURNING "replayID", "replayTitle", "stageName", "createdAt"
`

func (q *ReplayQueries) Create(ctx context.Context, arg CreateReplayParams) (Replay, error) {
	row := q.db.QueryRow(ctx, createReplay,
		arg.ReplayTitle,
		arg.StageName,
		arg.ReplayFileName,
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

const getReplayFileName = `--getReplayFileName
SELECT
	"replayFileName"
FROM
	"replays"
WHERE
	"replayID" = $1
LIMIT 1;
`

func (q *ReplayQueries) GetReplayFileName(ctx context.Context, replayID int) (string, error) {
	row := q.db.QueryRow(ctx, getReplayFileName,
		replayID,
	)
	var retData string
	err := row.Scan(
		&retData,
	)
	return retData, err
}

const getReplayByID = `--name: getReplayByID
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
WHERE
	r."replayID" = $1
GROUP BY
    r."replayID",
    r."replayTitle",
    r."stageName",
    r."createdAt"
LIMIT
	1;
`

const initialListComment = `--name: getPaginatedComment
SELECT "commentID", "commentContent", "createdAt"
FROM "replayComments"
ORDER BY "createdAt" DESC
`

const paginatedListComment = `--name: getPaginatedComment
SELECT "commentID", "commentContent", "createdAt"
FROM "replayComments"
WHERE "createdAt" < $1
OR ("createdAt" = $1 AND "commentID" > $2)
ORDER BY "createdAt" DESC
LIMIT 20;
`

func (q *ReplayQueries) GetReplayByID(ctx context.Context, replayID int, lastID *int, lastCreatedAt *pgtype.Timestamp) (*ReplayDetail, error) {
	row := q.db.QueryRow(ctx, getReplayByID,
		replayID,
	)
	var ret ReplayDetail
	err := row.Scan(
		&ret.ReplayID,
		&ret.ReplayTitle,
		&ret.StageName,
		&ret.CreatedAt,
		&ret.Likes,
		&ret.CommentCount,
	)
	if err != nil {
		return nil, err
	}

	// Currently not going to paginate comment, just going to show all
	var rowComments pgx.Rows
	if lastID == nil || lastCreatedAt == nil {
		rowComments, err = q.db.Query(ctx, initialListComment)
		if err != nil {
			return nil, err
		}
	} else {
		rowComments, err = q.db.Query(ctx, paginatedListComment, *lastCreatedAt, *lastID)
		if err != nil {
			return nil, err
		}
	}

	defer rowComments.Close()
	for rowComments.Next() {
		var retComment replaycomments.ReplayComments
		if err := rowComments.Scan(
			&retComment.CommentID,
			&retComment.CommentContent,
			&retComment.CreatedAt,
		); err != nil {
			return nil, err
		}
		ret.Comments = append(ret.Comments, retComment)
	}
	return &ret, nil
}
