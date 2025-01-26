package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type XenditPayment struct{
	host string
	httpConector HttpConector
}

func NewXenditClient(httpConector HttpConector, host string) XenditPayment {
	return XenditPayment{
		httpConector: httpConector,
		host: host,
	}
}

type (
	XenditPaymentReqBody struct {
		Currency string 
		Amount int
		PaymentMethod PaymentMethod
		Metadata Metadata
	}

	PaymentMethod struct {
		PaymentMethodType string 
		Reusability string
		ReferenceID string
		VirtualAccount VirtualAccount
	}

	VirtualAccount struct {
		ChannelCode string
		ChannelProperties ChannelProperties
	}

	ChannelProperties struct{
		CustomerName string
	}

	Metadata struct {
		SKU string
	}
)
func (x *XenditPayment) SendPaymentRequest(ctx context.Context) (paymentID string, err error) {
	// TODO: inject the http client

	// call xendit PaymentReq API

		// construct request body
	
	// handle error response 

	// handle sucess response 

	requestBody := XenditPaymentReqBody{}
	_ = requestBody
	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(requestBody)
	if err != nil {
		return "", err
	}


	endpoint := fmt.Sprintf("%s%s",x.host, "/payment_requests")
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, reqBody)

	res, err := x.httpConector.Do(httpReq)
	if err != nil {
		return "", err
	}

	_ = res 
	return "", err
}