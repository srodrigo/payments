package payments

type UUID interface {
	GetNextUUID() string
}

type RandomUUID struct{}

type PaymentsRepository struct {
	payments []*Payment
	Uuid     UUID
}

func (uuid *RandomUUID) GetNextUUID() string {
	return ""
}

func NewPaymentsRepository() *PaymentsRepository {
	return &PaymentsRepository{
		payments: make([]*Payment, 0),
		Uuid:     &RandomUUID{},
	}
}

func (repository *PaymentsRepository) Save(payment *Payment) *Payment {
	newPayment := Payment(*payment)
	newPayment.Id = repository.Uuid.GetNextUUID()

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
