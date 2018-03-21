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
		Id: "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
	}

	repository.payments = append(repository.payments, newPayment)

	return newPayment
}
