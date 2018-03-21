package payments

type PaymentsRepository struct {
	payments []*Payment
}

func NewPaymentsRepository() *PaymentsRepository {
	return &PaymentsRepository{
		payments: make([]*Payment, 0),
	}
}

func (repository *PaymentsRepository) Save(payment *Payment) *Payment {
	newPayment := &Payment{
		Id: payment.Id,
	}

	repository.payments = append(repository.payments, newPayment)

	return newPayment
}
