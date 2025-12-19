package paymentHandler

import (
	"github.com/TewApirat/items-shop-ms/modules/payment/paymentUsecase"
)

type (
	paymentGrpcHandler struct {
		PaymentUsecase paymentUsecase.PaymentUsecaseService
		
	}
)

func NewPaymentGrpcHandler(PaymentUsecase paymentUsecase.PaymentUsecaseService)*paymentGrpcHandler{
	return &paymentGrpcHandler{PaymentUsecase}
}
