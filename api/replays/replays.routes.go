package replays

import (
	"Golang-Replay-REST/api"

	"github.com/gin-gonic/gin"
)

type ReplayRouter struct {
	ctrl ReplayController
}

func NewRouter(ctrl ReplayController) ReplayRouter {
	return ReplayRouter{ctrl: ctrl}
}

func (router *ReplayRouter) Routes(rg *gin.RouterGroup) {
	r := rg.Group("replays")
	r.POST("/", router.ctrl.CreateContact)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX) {

	query := NewReplayQuery(db)
	controller := NewController(query)
	router := NewRouter(*controller)

	router.Routes(rg)
}
