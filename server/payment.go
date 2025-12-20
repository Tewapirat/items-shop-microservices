package server


import(
		"github.com/TewApirat/items-shop-ms/modules/payment/paymentRepository"
		"github.com/TewApirat/items-shop-ms/modules/payment/paymentUsecase"
		"github.com/TewApirat/items-shop-ms/modules/payment/paymentHandler"
)

func (s *server) paymentService(){
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Check
	payment.GET("", s.healthCheckService)
}