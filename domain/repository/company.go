package repository

import "payconfig/domain/model"

type CompanyRepository interface {
	FindByID(id int64) (*model.Company, error)
}
