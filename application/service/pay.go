package service

import (
	"payconfig/application/dto"
	"payconfig/domain/model"
	"payconfig/domain/service"
)

type PaymentConfigAppService struct {
	domainService *service.PaymentConfigService
}

func NewPaymentConfigAppService(ds *service.PaymentConfigService) *PaymentConfigAppService {
	return &PaymentConfigAppService{domainService: ds}
}

func (s *PaymentConfigAppService) SetPaymentConfig(configDTO dto.PaymentConfigDTO) error {
	// Convert DTO to domain model
	config := &model.PaymentConfig{
		CompanyID: configDTO.CompanyID,
		Methods:   make([]model.PaymentMethod, len(configDTO.Methods)),
	}
	for i, methodDTO := range configDTO.Methods {
		config.Methods[i] = model.PaymentMethod{
			Type:   methodDTO.Type,
			Config: methodDTO.Config,
		}
	}
	return s.domainService.SetPaymentConfig(config)
}

func (s *PaymentConfigAppService) GetPaymentConfig(companyID int64) (*dto.PaymentConfigDTO, error) {
	config, err := s.domainService.GetPaymentConfig(companyID)
	if err != nil {
		return nil, err
	}
	// Convert domain model to DTO
	configDTO := &dto.PaymentConfigDTO{
		CompanyID: config.CompanyID,
		Methods:   make([]dto.PaymentMethodDTO, len(config.Methods)),
	}
	for i, method := range config.Methods {
		configDTO.Methods[i] = dto.PaymentMethodDTO{
			Type:   method.Type,
			Config: method.Config,
		}
	}
	return configDTO, nil
}
