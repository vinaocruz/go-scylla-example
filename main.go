package main

import (
	"github.com/vinaocruz/go-scylla-example/config"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/server"
)

var appServer server.Server

func main() {
	config := config.InitConfig()
	db := database.NewScyllaDB(config)

	appServer = server.NewGinServer(db)
	appServer.Start()
}
