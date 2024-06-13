package strategy

import "testing"

func TestPayment_Pay(t *testing.T) {
	var payment *Payment

	payment = NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	payment = NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
