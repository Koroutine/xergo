package xero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Payment struct {
	Invoice Invoice `json:"Invoice"`
	Account Account `json:"Account"`
	Date    string  `json:"Date"`
	Amount  float64 `json:"Amount"`
}

type CreatePaymentResponse struct {
	Payments []Payment `json:"Payments"`
}

func (c *XeroClient) CreatePayment(payment *Payment) (*Payment, error) {
	req := c.SetupBaseRequest(PUT, "/api.xro/2.0/Payments")

	body, err := json.Marshal(payment)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(body))
	req.ContentLength = int64(len(body))

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {

		body, _ := io.ReadAll(response.Body)

		return nil, fmt.Errorf("could not create payment unexpected status code: %d: %s", response.StatusCode, string(body))
	}

	body, err = io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Println("xero response: ", string(body))
	}

	var paymentResponse CreatePaymentResponse

	err = json.Unmarshal(body, &paymentResponse)

	if err != nil {
		return nil, err
	}

	if len(paymentResponse.Payments) == 0 {
		return nil, fmt.Errorf("no payment returned")
	}

	return &paymentResponse.Payments[0], nil
}
