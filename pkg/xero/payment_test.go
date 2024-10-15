package xero

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCreatePayment(t *testing.T) {

	payment := &Payment{
		Invoice: Invoice{
			InvoiceID: "my-invoice-id",
			LineItems: []LineItem{},
		},
		Account: Account{
			AccountID: "my-account-id",
		},
		Date:   time.Now().Format(time.DateOnly),
		Amount: 10,
	}

	body, err := json.Marshal(payment)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(string(body))

}
