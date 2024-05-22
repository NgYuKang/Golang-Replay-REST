package replays

import (
	"Golang-Replay-REST/api"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	r.POST("/", router.ctrl.Create)
	r.GET("/", router.ctrl.List)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX, uploader *s3manager.Uploader, downloader *s3manager.Downloader) {

	query := NewReplayQuery(db)
	controller := NewController(query, uploader, downloader)
	router := NewRouter(*controller)

	router.Routes(rg)
}
