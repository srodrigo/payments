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
	newPayment := Payment(*payment)
	newPayment.Id = "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43"

	repository.payments = append(repository.payments, &newPayment)

	return &newPayment
}

func (repository *PaymentsRepository) FindById(id string) *Payment {
	for i, _ := range repository.payments {
		payment := repository.payments[i]
		if payment.Id == id {
			return payment
		}
	}

	return &Payment{}
}

func (repository *PaymentsRepository) FindAll() []*Payment {
	paymentsCopy := make([]*Payment, len(repository.payments))
	copy(paymentsCopy, repository.payments)

	return paymentsCopy
}
