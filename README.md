# ssl-wireless-pgw-golang-sdk (Beta)
> SSLCOMMERZ is the first payment gateway in Bangladesh opening doors for merchants to receive payments on the internet via their online stores. Their customers will be able to buy products online using their credit cards as well as bank accounts.

## Initialize payment

``` go

import (
    "github.com/sagar290/ssl_wireless_pgw_golang_sdk
)
func main() {
    isLive := false
	ssl := ssl_wireless_pgw_golang_sdk.Init("store_id", "store_password", isLive)

	body := ssl_wireless_pgw_golang_sdk.structs.PaymentBody{
		TranId:          "123fgh",
		ShippingMethod:  "online",
		ProductName:     "Test",
		ProductCategory: "Demo",
		ProductProfile:  "non-physical-goods",
		TotalAmount:     250,
		Currency:        "BDT",
	}

	body.SetCustomerDetails(structs.CustomerDetails{
		Name:  "Sagar290",
		Phone: "01612345678",
		Email: "tes@domin.com",
		Add1:  "-",
	})

	body.SetShippingDetails(structs.ShippingDetails{
		ShippingMethod: "NO",
		NumOfItem:      1,
	})

	body.SetSuccessUrl("https://google.com")
	body.SetCancelUrl("https://sagardash.com")
	body.SetFailUrl("https://bornonit.com")
	body.SetIpnUrl("https://10minuteschool.com")

	var response ssl.structs.PaymentResponse
    
        // initialize the payment
	response, _ = ssl.MakePayment(body)

	fmt.Println(response.GatewayPageURL)
}
```

### Structs
- Body
  - PaymentBody
  - OrderValidateBody
  - TransactionQueryBySessionKeyBody
  - TransactionQueryByTransactionId
- Response
  - IpnResponse
  - OrderValidateResponse
  - TransactionQueryBySessionKeyResponse
  - TransactionQueryByTransactionIdResponse
  - InitiateRefundResponse
  - RefundQueryURLResponse

### Methods
- MakePayment
- ValidatePayment **(for ipn)**
- InitiateRefund
- RefundQuery
- TransactionQueryBySessionId
- TransactionQueryByTransactionId

# Contribution guideline
Make these things more useful by contributing our brain and efforts. History will remember our contribution as far as lazy peoples like me are alive ;)

- Clone the repository
- Create branch from master
- make changes
- push the branch