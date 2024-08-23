package model

type PaymentConfig struct {
	ID        int64
	CompanyID int64
	Methods   []PaymentMethod
}
