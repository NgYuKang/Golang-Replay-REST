package replaycomments

import (
	"Golang-Replay-REST/api"

	"github.com/gin-gonic/gin"
)

type RepayCommentsRouter struct {
	ctrl ReplayCommentsController
}

func NewRouter(ctrl ReplayCommentsController) RepayCommentsRouter {
	return RepayCommentsRouter{ctrl: ctrl}
}

func (router *RepayCommentsRouter) Routes(rg *gin.RouterGroup) {
	r := rg.Group("replay-comments")
	r.POST("/", router.ctrl.Create)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX) {

	query := NewReplayCommentsQuery(db)
	controller := NewController(query)
	router := NewRouter(*controller)

	router.Routes(rg)
}
