package paymentHandler

import (
	"github.com/TewApirat/items-shop-ms/config"
	"github.com/TewApirat/items-shop-ms/modules/payment/paymentUsecase"
)

type (
		PaymentHttpHandlerService interface{}

		paymentHttpHandler struct {
			cfg 			*config.Config
			paymentUsecase paymentUsecase.PaymentUsecaseService
		}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService)PaymentHttpHandlerService{
	return &paymentHttpHandler{cfg, paymentUsecase}
}