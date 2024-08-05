package server

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/docs"
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
	driversRouters := r.Group("/v1/drivers")
	{
		handlers.NewDriverHandler(s.db, driversRouters)
	}
	vehiclesRouters := r.Group("/v1/drivers/:cnh/vehicles")
	{
		handlers.NewVehicleHandler(s.db, vehiclesRouters)
	}

	r.GET("/v1/ping", s.ping)

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}

// PingExample godoc
// @Summary ping example
// @Description do ping
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ping [get]
func (s *GinServer) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
