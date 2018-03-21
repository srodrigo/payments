package payments

type Payments struct {
	PaymentsRepository *PaymentsRepository
}

func NewPayments(paymentsRepository *PaymentsRepository) *Payments {
	return &Payments{paymentsRepository}
}

func (payments Payments) CreatePayment(payment Payment) {
}
