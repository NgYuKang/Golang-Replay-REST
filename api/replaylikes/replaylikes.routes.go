package replaylikes

import (
	"Golang-Replay-REST/api"

	"github.com/gin-gonic/gin"
)

type ReplayLikesRouter struct {
	ctrl ReplayLikesController
}

func NewRouter(ctrl ReplayLikesController) ReplayLikesRouter {
	return ReplayLikesRouter{ctrl: ctrl}
}

func (router *ReplayLikesRouter) Routes(rg *gin.RouterGroup) {
	r := rg.Group("replay-likes")
	r.POST("/", router.ctrl.Create)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX) {

	query := NewReplayLikesQuery(db)
	controller := NewController(query)
	router := NewRouter(*controller)

	router.Routes(rg)
}
