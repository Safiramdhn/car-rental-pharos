package controllers

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DriverController struct {
	service services.DriverService
	log     *zap.Logger
}

func NewDriverController(service services.DriverService, log *zap.Logger) *DriverController {
	return &DriverController{service: service, log: log}
}

func (ctl *DriverController) CreateDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, "Error binding JSON", 400)
		return
	}

	err := ctl.service.Create(&driver)
	if err != nil {
		ctl.log.Error("Error creating driver", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Created driver", zap.Uint("id", driver.ID))
	helpers.GoodResponseWithData(c, "Driver has been created", 201, nil)
}

func (ctl *DriverController) UpdateDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 400)
		return
	}
	idParam := c.Param("id")
	id, err := helpers.StringToUint(idParam)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	driver.ID = id
	err = ctl.service.Update(&driver)
	if err != nil {
		ctl.log.Error("Error updating driver", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Updated driver", zap.Uint("id", driver.ID))
	helpers.GoodResponseWithData(c, "Driver has been updated", 200, driver)
}

func (ctl *DriverController) DeleteDriver(c *gin.Context) {
	id := c.Param("id")
	driverID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	err = ctl.service.Delete(driverID)
	if err != nil {
		ctl.log.Error("Error deleting driver", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}
	ctl.log.Info("Deleted driver", zap.String("id", id))
	helpers.GoodResponseWithData(c, "Driver has been deleted", 200, nil)
}

func (ctl *DriverController) GetDriver(c *gin.Context) {
	id := c.Param("id")
	driverID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	driver, err := ctl.service.Get(driverID)
	if err != nil {
		ctl.log.Error("Error getting driver", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Get driver", zap.String("id", id))
	helpers.GoodResponseWithData(c, "Driver found", 200, driver)
}

func (ctl *DriverController) GetAllDrivers(c *gin.Context) {
	limit, _ := helpers.StringToUint(c.DefaultQuery("per_page", "10"))
	page, _ := helpers.StringToUint(c.DefaultQuery("page", "1"))
	drivers, countData, err := ctl.service.GetAll(limit, page)
	if err != nil {
		ctl.log.Error("Error getting drivers", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if drivers == nil {
		ctl.log.Info("No drivers found")
		helpers.GoodResponseWithPage(c, "No drivers found", 200, 0, 0, 0, 0, nil)
		return
	}

	ctl.log.Info("Get all drivers", zap.Uint("page", page), zap.Uint("per_page", limit))
	totalPage := (countData + int(limit) - 1) / int(limit) // calculate total pages correctly
	helpers.GoodResponseWithPage(c, "Drivers found", 200, countData, totalPage, int(page), int(limit), drivers)
}
