package controllers

import (
	"car-rental/helpers"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DriverIncentiveController struct {
	service services.DriverIncentiveService
	log     *zap.Logger
}

func NewDriverIncentiveController(service services.DriverIncentiveService, log *zap.Logger) *DriverIncentiveController {
	return &DriverIncentiveController{service: service, log: log}
}

func (ctl *DriverIncentiveController) GetAllDriverIncentives(c *gin.Context) {
	limit, _ := helpers.StringToUint(c.DefaultQuery("per_page", "10"))
	page, _ := helpers.StringToUint(c.DefaultQuery("page", "1"))
	driverIncentives, countData, err := ctl.service.GetAll(limit, page)
	if err != nil {
		ctl.log.Error("Error getting driverIncentives", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if driverIncentives == nil {
		ctl.log.Info("No driverIncentives found")
		helpers.GoodResponseWithPage(c, "No driverIncentives found", 200, 0, 0, 0, 0, nil)
		return
	}

	ctl.log.Info("Get all driverIncentives", zap.Uint("page", page), zap.Uint("per_page", limit))
	totalPage := (countData + int(limit) - 1) / int(limit) // calculate total pages correctly
	helpers.GoodResponseWithPage(c, "DriverIncentives found", 200, countData, totalPage, int(page), int(limit), driverIncentives)
}
