package payments

type PaymentsService struct {
	PaymentsRepository *PaymentsRepository
}

func NewPaymentsService(paymentsRepository *PaymentsRepository) *PaymentsService {
	return &PaymentsService{paymentsRepository}
}

func (service *PaymentsService) CreatePayment(payment *Payment) *Payment {
	return service.PaymentsRepository.Save(payment)
}
