package service

import (
	"payconfig/domain/model"
	"payconfig/domain/repository"
)

type PaymentConfigService struct {
	paymentConfigRepo repository.PaymentConfigRepository
	companyRepo       repository.CompanyRepository
}

func NewPaymentConfigService(pcr repository.PaymentConfigRepository, cr repository.CompanyRepository) *PaymentConfigService {
	return &PaymentConfigService{
		paymentConfigRepo: pcr,
		companyRepo:       cr,
	}
}

func (s *PaymentConfigService) SetPaymentConfig(config *model.PaymentConfig) error {
	return s.paymentConfigRepo.Save(config)
}

func (s *PaymentConfigService) GetPaymentConfig(companyID int64) (*model.PaymentConfig, error) {
	return s.paymentConfigRepo.FindByCompanyID(companyID)
}
