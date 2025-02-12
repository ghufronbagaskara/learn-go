package vo

type PaymentRequest struct {

}

type (
	XenditPaymentRequest struct {
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



	// requestBody := XenditPaymentReqBody{
	// 	Currency: "IDR",
	// 	Amount: 100000,
	// 	PaymentMethod: PaymentMethod{
	// 		PaymentMethodType: "VIRTUAL_ACCOUNT",
	// 		Reusability: "ONE_TIME",
	// 		ReferenceID: fmt.Sprintf("pm-level-%s", uuid.New().String()),
	// 		VirtualAccount: VirtualAccount{
	// 			ChannelCode: "BRI",
	// 			ChannelProperties: ChannelProperties{
	// 				CustomerName: "John Doe",
	// 			},
	// 		},
	// 	},
	// 	// Metadata: Metadata{
	// 	// 	SKU: "ABCDEFGH",
	// 	// },
	// }
	// _ = requestBody