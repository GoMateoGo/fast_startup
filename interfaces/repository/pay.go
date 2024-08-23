package repository

import (
	"payconfig/domain/model"
	"payconfig/domain/repository"
)

type PaymentConfigRepositoryImpl struct{}

func NewPaymentConfigRepositoryImpl() repository.PaymentConfigRepository {
	return &PaymentConfigRepositoryImpl{}
}

func (r *PaymentConfigRepositoryImpl) Save(config *model.PaymentConfig) error {
	//return dbmysql.DB.Update(config).Error
	return nil
}

func (r *PaymentConfigRepositoryImpl) FindByCompanyID(companyID int64) (*model.PaymentConfig, error) {
	var config model.PaymentConfig
	//result := persistence.DB.Where("company_id = ?", companyID).First(&config)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	return &config, nil
}
