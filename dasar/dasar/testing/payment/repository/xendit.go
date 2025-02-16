package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"unit-test/payment/vo"
)

type Config struct{
	AuthKey string
}

type XenditPayment struct{
	config Config
	host string
	httpConector HttpConector
}

func NewXenditClient(httpConector HttpConector, host string, authKey string) XenditPayment {
	return XenditPayment{
		httpConector: httpConector,
		host: host,
		config: Config{
			AuthKey: authKey,
		},
	}
}



func (x *XenditPayment) SendPaymentRequest(ctx context.Context, paymentRequest vo.XenditPaymentRequest) (paymentID string, err error) {
	if x.config.AuthKey == ""{
		return "", errors.New("empty auth key")
	}
	
	if paymentRequest.PaymentMethod.ReferenceID == "" {
		return "", errors.New("empty / invalid reference_id")
	}
	
	// TODO: inject the http client

	// validate request payload
	// call xendit PaymentReq API
	// construct request body
	// handle error response 
	// handle sucess response 

	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(paymentRequest)
	if err != nil {
		return "", err
	}


	endpoint := fmt.Sprintf("%s%s",x.host, "/payment_requests")
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, reqBody)
	httpReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", x.config.AuthKey))
	httpReq.Header.Set("Content-Type", "application/json")

	res, err := x.httpConector.Do(httpReq)
	if err != nil {
		return "", err
	}

	rawResponseBody, err := io.ReadAll(res.Body)
	strResponseBody := string(rawResponseBody)
	_ = strResponseBody

	_ = res 
	return "", err
}