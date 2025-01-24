package controllers

import (
	"car-rental/helpers"
	"car-rental/models"
	"car-rental/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CustomerController struct {
	service services.CustomerService
	log     *zap.Logger
}

func NewCustomerController(service services.CustomerService, log *zap.Logger) *CustomerController {
	return &CustomerController{service: service, log: log}
}

func (ctl *CustomerController) CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		ctl.log.Error("Error binding JSON", zap.Error(err))
		helpers.BadResponse(c, "Error binding JSON", 400)
		return
	}

	err := ctl.service.Create(&customer)
	if err != nil {
		ctl.log.Error("Error creating customer", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Created customer", zap.Uint("id", customer.ID))
	helpers.GoodResponseWithData(c, "Customer has been created", 201, nil)
}

func (ctl *CustomerController) UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
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
	customer.ID = id
	err = ctl.service.Update(&customer)
	if err != nil {
		ctl.log.Error("Error updating customer", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Updated customer", zap.Uint("id", customer.ID))
	helpers.GoodResponseWithData(c, "Customer has been updated", 200, customer)
}

func (ctl *CustomerController) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	customerID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	err = ctl.service.Delete(customerID)
	if err != nil {
		ctl.log.Error("Error deleting customer", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}
	ctl.log.Info("Deleted customer", zap.String("id", id))
	helpers.GoodResponseWithData(c, "Customer has been deleted", 200, nil)
}

func (ctl *CustomerController) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customerID, err := helpers.StringToUint(id)
	if err != nil {
		ctl.log.Error("Error converting string to uint", zap.Error(err))
		helpers.BadResponse(c, "Error converting string to uint", 400)
		return
	}
	customer, err := ctl.service.Get(customerID)
	if err != nil {
		ctl.log.Error("Error getting customer", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	ctl.log.Info("Get customer", zap.String("id", id))
	helpers.GoodResponseWithData(c, "Customer found", 200, customer)
}

func (ctl *CustomerController) GetAllCustomers(c *gin.Context) {
	limit, _ := helpers.StringToUint(c.DefaultQuery("per_page", "10"))
	page, _ := helpers.StringToUint(c.DefaultQuery("page", "1"))
	customers, countData, err := ctl.service.GetAll(limit, page)
	if err != nil {
		ctl.log.Error("Error getting customers", zap.Error(err))
		helpers.BadResponse(c, err.Error(), 500)
		return
	}

	if customers == nil {
		ctl.log.Info("No customers found")
		helpers.GoodResponseWithPage(c, "No customers found", 200, 0, 0, 0, 0, nil)
		return
	}

	ctl.log.Info("Get all customers", zap.Uint("page", page), zap.Uint("per_page", limit))
	totalPage := (countData + int(limit) - 1) / int(limit) // calculate total pages correctly
	helpers.GoodResponseWithPage(c, "Customers found", 200, countData, totalPage, int(page), int(limit), customers)
}
