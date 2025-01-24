package controllers

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CarController struct {
	service services.CarService
	log     *zap.Logger
}

func NewCarController(service services.CarService, log *zap.Logger) *CarController {
	return &CarController{service: service, log: log}
}

func (ctl *CarController) CreateCar(c *gin.Context) {
	// Get the request body
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, "Error binding JSON", 400)
		return
	}
	// Create the car
	err := ctl.service.Create(&car)
	if err != nil {
		ctl.log.Error("Error creating car", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}
	// Return success response
	ctl.log.Info("Created car", zap.Uint("id", car.ID))
	helpers.GoodResponseWithData(c, "Car has been created", 201, nil)
}

func (ctl *CarController) UpdateCar(c *gin.Context) {
	// Get the request body
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 400)
		return
	}
	// Get the car ID
	idParam := c.Param("id")
	id, err := helpers.StringToUint(idParam)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	car.ID = id
	// Update the car
	err = ctl.service.Update(&car)
	if err != nil {
		ctl.log.Error("Error updating car", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}
	// Return success response
	ctl.log.Info("Updated car", zap.Uint("id", car.ID))
	helpers.GoodResponseWithData(c, "Car has been updated", 200, car)
}

func (ctl *CarController) DeleteCar(c *gin.Context) {
	// Get the car ID
	id := c.Param("id")
	carID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	// Delete the car
	err = ctl.service.Delete(carID)
	if err != nil {
		ctl.log.Error("Error deleting car", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}
	// Return success response
	ctl.log.Info("Deleted car", zap.Uint("id", carID))
	helpers.GoodResponseWithData(c, "Car has been deleted", 200, nil)
}

func (ctl *CarController) GetCar(c *gin.Context) {
	// Get the car ID
	id := c.Param("id")
	carID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}

	// Get the car
	car, err := ctl.service.Get(carID)
	if err != nil {
		ctl.log.Error("Error getting car", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	// Return success response
	ctl.log.Info("Get car", zap.String("id", id))
	helpers.GoodResponseWithData(c, "Car found", 200, car)
}

func (ctl *CarController) GetAllCars(c *gin.Context) {
	// get limit and page parameters
	limit, _ := helpers.StringToUint(c.DefaultQuery("per_page", "10"))
	page, _ := helpers.StringToUint(c.DefaultQuery("page", "1"))

	// Get all cars
	cars, countData, err := ctl.service.GetAll(limit, page)
	if err != nil {
		ctl.log.Error("Error getting cars", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if cars == nil {
		ctl.log.Info("No cars found")
		helpers.GoodResponseWithPage(c, "No cars found", 200, 0, 0, 0, 0, nil)
	}

	// Return success response
	ctl.log.Info("Get all cars", zap.Int("count", countData))
	totalPage := (countData + int(limit) - 1) / int(limit) // calculate total pages correctly
	helpers.GoodResponseWithPage(c, "Cars found", 200, countData, totalPage, int(page), int(limit), cars)
}
