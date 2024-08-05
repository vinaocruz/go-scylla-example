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

// @Summary recupera veiculo
// @Description Este endpoint busca um veiculo especifico de um motorista através de sua placa
// @Tags Veiculos
// @Accept json
// @Produce json
// @Success 200 {object} entities.Driver
// @failure 404 {string} string
// @Router /drivers/{cnh}/vehicles/{license_plate} [get]
// @Param cnh   		path string true "Número de CNH"
// @Param license_plate path string true "Número da Placa do Veiculo"
func (handler *VehicleHandler) GetVehicle(c *gin.Context) {
	driver, err := handler.repository.Load(c.Param("cnh"), c.Param("license_plate"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, driver)
}

// @Summary edita veiculo
// @Description Este endpoint permite editar dados de um veiculo especifico de um motorista
// @Tags Veiculos
// @Accept json
// @Produce json
// @Success 200 {object} entities.Driver
// @failure 404 {string} string
// @failure 400 {string} string "failed to bind json"
// @failure 500 {string} string "failed to update"
// @Router /drivers/{cnh}/vehicles/{license_plate} [put]
// @Param cnh   		path string true "Número de CNH"
// @Param license_plate path string true "Número da Placa do Veiculo"
// @Param model body string true "Modelo do veiculo"
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
		c.JSON(http.StatusBadRequest, gin.H{
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

	c.JSON(http.StatusOK, driver)
}

// @Summary remove veiculo
// @Description Este endpoint remove veiculo especifico de um motorista
// @Tags Veiculos
// @Accept json
// @Produce json
// @Success 204
// @failure 404 {string} string
// @failure 500 {string} string "failed to delete"
// @Router /drivers/{cnh}/vehicles/{license_plate} [delete]
// @Param cnh   		path string true "Número de CNH"
// @Param license_plate path string true "Número da Placa do Veiculo"
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
