package company

import (
	"net/http"
	"payconfig/application/dto"
	"payconfig/application/service"
	dservice "payconfig/domain/service"
	"payconfig/interfaces/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentConfigHandler struct {
	appService *service.PaymentConfigAppService
}

func NewPaymentConfigHandler(appService *service.PaymentConfigAppService) *PaymentConfigHandler {
	return &PaymentConfigHandler{appService: appService}
}

func (h *PaymentConfigHandler) SetPaymentConfig(c *gin.Context) {
	var configDTO dto.PaymentConfigDTO
	if err := c.ShouldBindJSON(&configDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.appService.SetPaymentConfig(configDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment configuration set successfully"})
}

func (h *PaymentConfigHandler) GetPaymentConfig(c *gin.Context) {
	companyIDStr := c.Query("company_id")
	companyID, err := strconv.ParseInt(companyIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	config, err := h.appService.GetPaymentConfig(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

func SetupRoutes(r *gin.Engine) {
	handler := NewPaymentConfigHandler(service.NewPaymentConfigAppService(
		dservice.NewPaymentConfigService(
			repository.NewPaymentConfigRepositoryImpl(),
			repository.NewCompanyRepositoryImpl(),
		),
	))

	r.POST("/payment-config", handler.SetPaymentConfig)
	r.GET("/payment-config", handler.GetPaymentConfig)
}
