package services

import (
	"car-rental/models"
	"car-rental/repositories"

	"go.uber.org/zap"
)

type MembershipService interface {
	UpdateCustomerMembership(input models.MembershipDTO) (*models.Customer, error)
}

type membershipService struct {
	repo repositories.Repository
	log  *zap.Logger
}

// UpdateCustomerMembership implements MembershipService.
func (m *membershipService) UpdateCustomerMembership(input models.MembershipDTO) (*models.Customer, error) {
	customer, err := m.repo.Customer.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	customer.MembershipID = &input.MembershipID
	if *customer.MembershipID == 0 {
		customer.MembershipID = nil
	}

	m.log.Info("Updating customer membership", zap.Uint("ID", input.ID))
	err = m.repo.Customer.Update(customer)
	return customer, err
}

func NewMembershipService(repo repositories.Repository, log *zap.Logger) MembershipService {
	return &membershipService{
		repo: repo,
		log:  log,
	}
}
