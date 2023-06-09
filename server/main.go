package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/config"
	"github.com/qanyue/aldb/server/middleware"
	"github.com/qanyue/aldb/server/router"
)

// @title Aldb Example API
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api

func main() {
	r := SetRouter()
	r.Run(fmt.Sprintf(":%v", config.Conf.Port))
}
func SetRouter() *gin.Engine {
	// Init Logger
	if err := middleware.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Println(err)
		return nil
	}
	r := router.InitRouter()

	//r.Run(fmt.Sprintf(":%v", config.Conf.Port))
	return r
}
