package service

type PaymentGatewayProvider interface {
	SendPaymentRequest() (paymentID string, err error)
}

type Payment struct {
	xendit PaymentGatewayProvider
}

func (p *Payment) Pay() error {
	// create transactionID
	// insert into DB: postgress
	// construct payment request payload
	// run request payment validation logic
	// call third party API 
	_, err := p.xendit.SendPaymentRequest()
	if err != nil {
		return err 
	}

	return nil
}