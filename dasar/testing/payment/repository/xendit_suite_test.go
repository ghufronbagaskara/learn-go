package repository_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"unit-test/payment/repository"
	"unit-test/payment/repository/mock"
	"unit-test/payment/vo"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

// using test suite
// SetupSuite() - before test
// SetupTest() - before test cases
// TearDownTest() - after each test cases
// TearDownSuite() - after test

type XendiPaymentTestSuite struct {
	suite.Suite
	ctx context.Context
	httpClientMock *mock.MockHttpConector 
	xenditCLient repository.XenditPayment
	xenditHost string
	xenditAuthKey string
}

func TestXenditPayment (t *testing.T){
	suite.Run(t, new(XendiPaymentTestSuite))
}


func (s *XendiPaymentTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	httpClientMock := mock.NewMockHttpConector(ctrl)
	s.xenditHost = "http://mock.server"
	s.xenditAuthKey = "supersecret"
	s.ctx = context.Background() 

	xenditClient := repository.NewXenditClient(httpClientMock, s.xenditHost, s.xenditAuthKey)

	s.httpClientMock = httpClientMock
	s.xenditCLient = xenditClient
}

func (s *XendiPaymentTestSuite) SetupTest() {
	
}

func (s *XendiPaymentTestSuite) TestXenditPayment_SendPaymentRequest_EmptyAuthHeader() {
	httpClientMock := mock.NewMockHttpConector(gomock.NewController(s.T()))
	host := "http://mock.server"
	authKey := ""

	xenditClient := repository.NewXenditClient(httpClientMock, host, authKey)
	ctx := context.Background()
	_, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})

	s.Error(err, "it should return error due to empty auth key ")
}

func (s *XendiPaymentTestSuite) TestXenditPayment_SendPaymentRequest_WithEmptyPayload() {
	_, gotErr := s.xenditCLient.SendPaymentRequest(s.ctx, vo.XenditPaymentRequest{})
	s.Error(gotErr, "it should return error due to empty payment request payload")
}

func (s *XendiPaymentTestSuite) TestXenditPayment_SendPaymentRequest_ErrorWhileSendingHtppRequest() {
	s.httpClientMock.EXPECT().Do(gomock.Any()).Return(&http.Response{}, errors.New("Http failure"))
	paymentRequest := vo.XenditPaymentRequest{
		Currency: "IDR",
		PaymentMethod: vo.PaymentMethod{
			PaymentMethodType: "VIRTUAL_ACCOUNT",
			ReferenceID: "random-id",
		},
	}
	_, gotErr := s.xenditCLient.SendPaymentRequest(s.ctx, paymentRequest)
	s.Error(gotErr, "it should return error due to http failure")
	s.ErrorContains(gotErr, "failure") 
}

