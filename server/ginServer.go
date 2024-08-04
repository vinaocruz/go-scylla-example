package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/infrastructure/handlers"
)

type GinServer struct {
	db *database.Database
}

func NewGinServer(db *database.Database) *GinServer {
	return &GinServer{
		db: db,
	}
}

func (s *GinServer) Start() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	driversRouters := r.Group("/drivers")
	{
		handlers.NewDriverHandler(s.db, driversRouters)
	}

	vehiclesRouters := r.Group("/drivers/:cnh/vehicles")
	{
		handlers.NewVehicleHandler(s.db, vehiclesRouters)
	}

	r.Run()
}
