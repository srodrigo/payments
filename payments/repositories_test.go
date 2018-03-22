package payments

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSavesPayments(t *testing.T) {
	paymentsRepository := NewPaymentsRepository()

	paymentsRepository.Save(&Payment{})
	paymentsRepository.Save(&Payment{})

	allPayments := paymentsRepository.FindAll()
	assert.Equal(t, 2, len(allPayments))
}

func TestSavesPaymentWithRandomId(t *testing.T) {
	paymentsRepository := NewPaymentsRepository()
	hexadigit := "[0-9a-f]"
	r, _ := regexp.Compile(fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$",
		hexadigit, hexadigit, hexadigit, hexadigit, hexadigit))

	savedPayment := paymentsRepository.Save(&Payment{})

	assert.True(t, r.MatchString(savedPayment.Id))
}

func TestFindsPaymentsById(t *testing.T) {
	paymentsRepository := NewPaymentsRepository()
	payment := paymentsRepository.Save(&Payment{})

	paymentById, err := paymentsRepository.FindById(payment.Id)

	assert.Equal(t, payment, paymentById)
	assert.Equal(t, nil, err)
}
