package internal

import (
	"fmt"
	"net/http"

	"github.com/anhcx0209/knovel-hometest/config"
	"github.com/anhcx0209/knovel-hometest/database"
	internal "github.com/anhcx0209/knovel-hometest/internal/http"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	engine *gin.Engine
	db     database.Database
	conf   *config.Config
}

func NewGinServer(conf *config.Config, db database.Database) internal.Server {
	engine := gin.Default()
	return &ginServer{
		engine: engine,
		db:     db,
		conf:   conf,
	}
}

func (s *ginServer) Start() {
	s.engine.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	s.engine.Run(fmt.Sprintf(":%d", s.conf.Server.Port))
}
