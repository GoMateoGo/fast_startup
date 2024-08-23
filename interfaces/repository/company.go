package repository

import (
	dbmysql "payconfig/core/db"
	"payconfig/domain/model"
	"payconfig/domain/repository"
)

type CompanyRepositoryImpl struct{}

func NewCompanyRepositoryImpl() repository.CompanyRepository {
	return &CompanyRepositoryImpl{}
}

func (r *CompanyRepositoryImpl) FindByID(id int64) (*model.Company, error) {
	var company model.Company
	company.ID = id
	_, err := dbmysql.DB.Get(&company)
	if err != nil {
		return nil, err
	}
	return &company, nil
}
