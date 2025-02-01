package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
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

type (
	XenditPaymentReqBody struct {
		Currency string `json:"currency"`
		Amount int 		`json:"amount"`
		PaymentMethod PaymentMethod `json:"payment_method"`
		Metadata Metadata `json:"metadata"`
	}

	PaymentMethod struct {
		PaymentMethodType string `json:"type"`
		Reusability string `json:"reusability"`
		ReferenceID string `json:"reference_id"`
		VirtualAccount VirtualAccount `json:"virtual_account"`
	}

	VirtualAccount struct {
		ChannelCode string `json:"channel_code"`
		ChannelProperties ChannelProperties `json:"channel_properties"`
	}

	ChannelProperties struct{
		CustomerName string `json:"customer_name"`
	}

	Metadata struct {
		SKU string `json:"sku"`
	}
)
func (x *XenditPayment) SendPaymentRequest(ctx context.Context) (paymentID string, err error) {
	// TODO: inject the http client

	// call xendit PaymentReq API

		// construct request body
	
	// handle error response 

	// handle sucess response 

	requestBody := XenditPaymentReqBody{
		Currency: "IDR",
		Amount: 100000,
		PaymentMethod: PaymentMethod{
			PaymentMethodType: "VIRTUAL_ACCOUNT",
			Reusability: "ONE_TIME",
			ReferenceID: fmt.Sprintf("pm-level-%s", uuid.New().String()),
			VirtualAccount: VirtualAccount{
				ChannelCode: "BRI",
				ChannelProperties: ChannelProperties{
					CustomerName: "John Doe",
				},
			},
		},
		Metadata: Metadata{
			SKU: "ABCDEFGH",
		},
	}
	_ = requestBody
	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(requestBody)
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