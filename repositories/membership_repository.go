package repositories

import (
	"car-rental/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MembershipRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewMembershipRepository(db *gorm.DB, log *zap.Logger) *MembershipRepository {
	return &MembershipRepository{db: db, log: log}
}

func (repo *MembershipRepository) FindByID(id uint) (*models.Membership, error) {
	var membership models.Membership
	err := repo.db.First(&membership, id).Error
	if err != nil {
		repo.log.Error("Error finding membership", zap.Error(err))
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("membership not found")
		}
		return nil, err
	}
	return &membership, nil
}
