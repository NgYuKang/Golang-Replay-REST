package replays

import (
	"Golang-Replay-REST/api"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dutchcoders/go-clamd"
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
	r.GET("/file/:replayID", router.ctrl.DownloadReplay)
}

func SetupRoutes(rg *gin.RouterGroup, db api.DBTX, uploader *s3manager.Uploader, downloader *s3manager.Downloader, cav *clamd.Clamd) {

	query := NewReplayQuery(db)
	controller := NewController(query, uploader, downloader, cav)
	router := NewRouter(*controller)

	router.Routes(rg)
}
