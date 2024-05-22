package replaycomments

type ReplayComments struct {
	ReplayID       int    `json:"replayID"`
	CommentID      int    `json:"commentID"`
	CommentContent string `json:"commentContent"`
}
