package repository

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"unit-test/payment/repository/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Negative Test Case
// sending request with empty auth header
// sending request with invalid auth header
// sending request with valid auth header
// sending request with empty request body
// sending request with broken request body
// sending request with incomplete
// sending request with invalid request body (missing some request field)
// sending request with duplicate referenceID
// sending request with inactive channel_code: BCA

// Positive Test Case
// sending with complete and valid request body

// Edge Case
// TODO: find the edge cases


func TestXenditPayment_SendPaymentRequest_WithEmptyAuthHeader(t *testing.T){
	httpClient := &http.Client{}
	hostName := "https://api.xendit.co"

	xenditClient := NewXenditClient(httpClient, hostName)
	ctx := context.Background()
	paymentID, err := xenditClient.SendPaymentRequest(ctx)
	if err != nil {
		t.Fatalf("it should not return error, but got %s", err.Error())
	}

	if paymentID == ""{
		t.Errorf("it should return empty paymentID %s", paymentID)
	}
}
	
func TestXenditPayment_SendPaymentRequest_IncompletedRequestData(t *testing.T){
	
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGot500(t *testing.T){

}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_Got500_ButEmptyResponseBody(t *testing.T){
	
}

func TestXenditPayment_SendPaymentRequest_CompleteRequestData_ButGotBrokenResponseBody(t *testing.T){
	
}

func TestXenditPayment_SendPaymentRequest_SuccessResponse(t *testing.T){
	
}

func TestXenditPayment_SendPaymentRequest(t *testing.T){
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
	host := "http://mock.server"

	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))
	
	xenditClient := NewXenditClient(httpClientMock, host)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
	assert.Error(t, err, "it should not return error")

	assert.Empty(t, paymentID, "it should return a valid created paymentID")
}