package controllers

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MembershipController struct {
	service services.MembershipService
	log     *zap.Logger
}

func NewMembershipController(service services.Service, log *zap.Logger) *MembershipController {
	return &MembershipController{service: service.Membership, log: log}
}

func (ctl *MembershipController) SetMembership(c *gin.Context) {
	// get customer id
	customerIdParam := c.Param("customer_id")
	customerId, err := helpers.StringToUint(customerIdParam)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}

	// get membership type id
	var membership models.MembershipDTO
	if err := c.ShouldBindJSON(&membership); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 400)
		return
	}

	membership.ID = customerId
	customer, err := ctl.service.UpdateCustomerMembership(membership)
	if err != nil {
		ctl.log.Error("Error updating customer", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Customer memberships updated", zap.Uint("ID", customer.ID), zap.Uint("MembershipID", *customer.MembershipID))
	helpers.GoodResponseWithData(c, "Customer memberships updated", 200, customer)
}
