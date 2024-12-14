package main

import (
	"github.com/anhcx0209/knovel-hometest/config"
	"github.com/anhcx0209/knovel-hometest/database"
	internal "github.com/anhcx0209/knovel-hometest/internal/http/gin"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	// server := internal.NewEchoServer(conf, db)
	server := internal.NewGinServer(conf, db)
	server.Start()
}
