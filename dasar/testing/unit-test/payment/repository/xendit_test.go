package repository

import (
	"context"
	"errors"
	"testing"
	"unit-test/payment/repository/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestXenditPayment_SendPaymentRequest(t *testing.T){
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
	host := "http://mock.server"

	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))
	
	xenditClient := NewXenditClient(httpClientMock, host)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
	assert.Error(t, err, "it should not return error")

	assert.Empty(t, paymentID, "it should return a valid created paymentID")
}