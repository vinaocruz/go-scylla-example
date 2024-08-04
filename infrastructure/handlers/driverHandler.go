package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/domain/entities"
	domain "github.com/vinaocruz/go-scylla-example/domain/repositories"
	infra "github.com/vinaocruz/go-scylla-example/infrastructure/repositories"
)

type DriverHandler struct {
	repository domain.DriverRepository
}

func NewDriverHandler(db *database.Database, router *gin.RouterGroup) {
	handler := &DriverHandler{
		repository: infra.NewScyllaDriverRepository(db),
	}

	router.GET("/", handler.ListDrivers)
	router.POST("/", handler.CreateDriver)
	router.GET("/:cnh", handler.GetDriver)
	router.PUT("/:cnh", handler.UpdateDriver)
	router.DELETE("/:cnh", handler.DeleteDriver)
}

func (handler *DriverHandler) ListDrivers(c *gin.Context) {
	listAll, err := handler.repository.ListAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list all",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": listAll,
	})
}

func (handler *DriverHandler) CreateDriver(c *gin.Context) {
	var driver entities.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to bind json",
		})
		return
	}

	driver.ID = gocql.TimeUUID()
	driver.CreatedAt = time.Now()

	err := handler.repository.Store(driver)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to store",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": driver,
	})
}

func (handler *DriverHandler) GetDriver(c *gin.Context) {
	drivers, err := handler.repository.Load(c.Param("cnh"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": drivers,
	})
}

func (handler *DriverHandler) UpdateDriver(c *gin.Context) {
	cnh := c.Param("cnh")

	drivers, err := handler.repository.Load(cnh)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var people entities.People

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to bind json",
		})
		return
	}

	for i := range drivers {
		drivers[i].Name = people.Name
		err = handler.repository.Store(drivers[i])

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": drivers,
	})
}

func (handler *DriverHandler) DeleteDriver(c *gin.Context) {
	err := handler.repository.Delete(c.Param("cnh"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
