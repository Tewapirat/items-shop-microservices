package paymentHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/payment/paymentUsecase"
)

type (
		PaymentQueueHandlerService interface{}

		paymentQueueHandler struct {
			cfg 			*config.Config
			paymentUsecase paymentUsecase.PaymentUsecaseService
		}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService)PaymentQueueHandlerService{
	return &paymentQueueHandler{cfg, paymentUsecase}
}