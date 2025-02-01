package repository

import (
	"testing"
)

// Negative Test Case
// sending request with empty auth header
// sending request with invalid auth header
// sending request with valid auth header
// sending request with empty request body:
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"amount must be greater than 0\"}\n"
// "{\"error_code\":\"API_VALIDATION_ERROR\",\"message\":\"Only one of 'payment_method' or 'payment_method_id' should be present per request\"}\n"
// sending request with broken request body
// sending request with incomplete
// sending request with invalid request body (missing some request field)
// sending request with duplicate referenceID

// Positive Test Case
// "{\"id\":\"pr-1a12c5ec-cbfc-45bd-ad0b-51e699867ac1\",\"country\":\"ID\",\"amount\":100000,\"currency\":\"IDR\",\"business_id\":\"599bd7f1ccab55b020bb1147\",\"reference_id\":\"0d4cd7f0-30ef-45b0-a756-7f2194f04b3a\",\"payment_method\":{\"id\":\"pm-dcc65478-e017-45f0-b772-714c709fafac\",\"type\":\"VIRTUAL_ACCOUNT\",\"reference_id\":\"pm-level-c1a883c7-5e07-4024-ad74-3dc728341d2e\",\"description\":null,\"created\":\"2025-02-01T08:11:43.826892322Z\",\"updated\":\"2025-02-01T08:11:44.06893992Z\",\"card\":null,\"ewallet\":null,\"direct_debit\":null,\"direct_bank_transfer\":null,\"over_the_counter\":null,\"virtual_account\":{\"amount\":100000,\"currency\":\"IDR\",\"channel_code\":\"BRI\",\"channel_properties\":{\"customer_name\":\"John Doe\",\"virtual_account_number\":\"262158018439804\",\"expires_at\":\"2056-02-01T08:11:43.883582Z\"}},\"qr_code\":null,\"metadata\":null,\"billing_information\":{\"city\":null,\"country\":\"\",\"postal_code\":null,\"province_state\":null,\"street_line1\":null,\"street_line2\":null},\"reusability\":\"ONE_TIME_USE\",\"status\":\"PENDING\"},\"description\":null,\"metadata\":{\"sku\":\"ABCDEFGH\"},\"customer_id\":null,\"capture_method\":\"AUTOMATIC\",\"initiator\":null,\"card_verification_results\":null,\"created\":\"2025-02-01T08:11:43.772237142Z\",\"updated\":\"2025-02-01T08:11:43.772237142Z\",\"status\":\"PENDING\",\"actions\":[],\"failure_code\":null,\"channel_properties\":null,\"shipping_information\":null,\"items\":null}"

// sending with complete and valid request body
// sending with inactive channel_code : BCA

// Edge Case
// TODO: find the edge cases

// func TestXenditPayment_SendPaymentRequest_ApiExploration(t *testing.T){
// 	httpClient := &http.Client{}
// 	hostName := "https://api.xendit.co"
// 	// apiKey := "something"
// 	authKey := "eG5kX2RldmVsb3BtZW50X09vbUFmT1V0aCtHb3dzWTZMZUpPSHpMQ1p0U2o4NEo5a1hEbitSeGovbUhXK2J5aERRVnhoZz09Og=="

// 	xenditClient := NewXenditClient(httpClient, hostName, authKey)
// 	ctx := context.Background()
// 	paymentID, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})
// 	if err != nil {
// 		t.Fatalf("it should not return error, but got %s", err.Error())
// 	}

// 	if paymentID == ""{
// 		t.Errorf("it should return empty paymentID %s", paymentID)
// 	}
// }

// func TestXenditPayment_SendPaymentRequest_EmptyAuthHeader(t *testing.T){
// 	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
// 	host := "http://mock.server"
// 	authKey := "supersecret"

// 	xenditClient := NewXenditClient(httpClientMock, host, authKey)
// 	ctx := context.Background()
// 	_, err := xenditClient.SendPaymentRequest(ctx, vo.XenditPaymentRequest{})

// 	if err == nil {
// 		t.Fatal("it should return error due to empty auth key")
// 	}
// }

// func TestXenditPayment_SendPaymentRequest_WithEmptyPayload(t *testing.T){
// 	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
// 	host := "http://mock.server"
// 	authKey := ""

// 	xenditClient := NewXenditClient(httpClientMock, host, authKey)
// 	ctx := context.Background()
// 	paymentReq := vo.XenditPaymentRequest{}

// 	_, err := xenditClient.SendPaymentRequest(ctx, paymentReq)

// 	if err == nil {
// 		t.Fatal("it should return error due to empty auth key")
// 	}
// }

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

// func TestXenditPayment_SendPaymentRequest(t *testing.T){
// 	httpClientMock := mock.NewMockHttpConector(gomock.NewController(t))
// 	host := "http://mock.server"

// 	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))
	
// 	xenditClient := NewXenditClient(httpClientMock, host)
// 	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
// 	assert.Error(t, err, "it should not return error")

// 	assert.Empty(t, paymentID, "it should return a valid created paymentID")
// }