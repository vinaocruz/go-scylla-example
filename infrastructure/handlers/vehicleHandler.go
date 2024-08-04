package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinaocruz/go-scylla-example/database"
	"github.com/vinaocruz/go-scylla-example/domain/entities"
	domain "github.com/vinaocruz/go-scylla-example/domain/repositories"
	infra "github.com/vinaocruz/go-scylla-example/infrastructure/repositories"
)

type VehicleHandler struct {
	repository       domain.VehicleRepository
	driverRepository domain.DriverRepository
}

func NewVehicleHandler(db *database.Database, router *gin.RouterGroup) {
	handler := &VehicleHandler{
		repository:       infra.NewScyllaVehicleRepository(db),
		driverRepository: infra.NewScyllaDriverRepository(db),
	}

	router.GET("/:license_plate", handler.GetVehicle)
	router.PUT("/:license_plate", handler.UpdateVehicle)
	router.DELETE("/:license_plate", handler.DeleteVehicle)
}

func (handler *VehicleHandler) GetVehicle(c *gin.Context) {
	driver, err := handler.repository.Load(c.Param("cnh"), c.Param("license_plate"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": driver,
	})
}

func (handler *VehicleHandler) UpdateVehicle(c *gin.Context) {
	cnh := c.Param("cnh")
	licensePlate := c.Param("license_plate")

	driver, err := handler.repository.Load(cnh, licensePlate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var vehicle entities.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to bind json",
		})
		return
	}

	driver.Model = vehicle.Model
	err = handler.driverRepository.Store(driver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": driver,
	})
}

func (handler *VehicleHandler) DeleteVehicle(c *gin.Context) {
	cnh := c.Param("cnh")
	licensePlate := c.Param("license_plate")

	driver, err := handler.repository.Load(cnh, licensePlate)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = handler.repository.Delete(driver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
