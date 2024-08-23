package repository

import "payconfig/domain/model"

type PaymentConfigRepository interface {
	Save(config *model.PaymentConfig) error
	FindByCompanyID(companyID int64) (*model.PaymentConfig, error)
}
