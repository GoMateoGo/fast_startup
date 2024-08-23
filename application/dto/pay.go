package dto

type PaymentConfigDTO struct {
	CompanyID int64              `json:"company_id"`
	Methods   []PaymentMethodDTO `json:"methods"`
}

type PaymentMethodDTO struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}
