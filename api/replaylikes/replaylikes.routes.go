package replaylikes

import (
	"Golang-Replay-REST/api"

	"github.com/gin-gonic/gin"
)

type ReplayRouter struct {
	ctrl ReplayLikesController
}

func NewRouter(ctrl ReplayLikesController) ReplayRouter {
	return ReplayRouter{ctrl: ctrl}
}

func (router *ReplayRouter) Routes(rg *gin.RouterGroup) {
	r := rg.Group("replay-likes")
	r.POST("/", router.ctrl.Create)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX) {

	query := NewReplayLikesQuery(db)
	controller := NewController(query)
	router := NewRouter(*controller)

	router.Routes(rg)
}
