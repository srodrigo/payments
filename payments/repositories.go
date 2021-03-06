package payments

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
)

type UUID interface {
	GetNextUUID() string
}

type RandomUUID struct{}

type PaymentsRepository struct {
	payments []*Payment
	Uuid     UUID
}

func (randomUuid *RandomUUID) GetNextUUID() string {
	return uuid.Must(uuid.NewV4()).String()
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

func (repository *PaymentsRepository) Update(id string, payment *Payment) (*Payment, error) {
	for i := 0; i < len(repository.payments); i++ {
		current := repository.payments[i]
		if current.Id == id {
			updatedPayment := Payment(*payment)
			updatedPayment.Id = current.Id
			repository.payments[i] = &updatedPayment
			return &updatedPayment, nil
		}
	}

	return &Payment{}, errors.New(fmt.Sprintf("Could not update payment with id %s, payment does not exist", id))
}

func (repository *PaymentsRepository) Delete(id string) error {
	for i := 0; i < len(repository.payments); i++ {
		current := repository.payments[i]
		if current.Id == id {
			repository.payments = append(repository.payments[:i], repository.payments[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Could not delete payment with id %s, payment does not exist", id))
}

func (repository *PaymentsRepository) FindById(id string) (*Payment, error) {
	for i := 0; i < len(repository.payments); i++ {
		current := repository.payments[i]
		if current.Id == id {
			return current, nil
		}
	}

	return &Payment{}, errors.New(fmt.Sprintf("Could not find payment with id %s", id))
}

func (repository *PaymentsRepository) FindAll() []*Payment {
	paymentsCopy := make([]*Payment, len(repository.payments))
	copy(paymentsCopy, repository.payments)

	return paymentsCopy
}
