package payments

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSavesPaymentWithRandomId(t *testing.T) {
	paymentsRepository := NewPaymentsRepository()
	hexadigit := "[0-9a-f]"
	r, _ := regexp.Compile(fmt.Sprintf("^%s{8}-%s{4}-%s{4}-%s{4}-%s{12}$",
		hexadigit, hexadigit, hexadigit, hexadigit, hexadigit))

	savedPayment := paymentsRepository.Save(&Payment{})

	assert.True(t, r.MatchString(savedPayment.Id))
}
