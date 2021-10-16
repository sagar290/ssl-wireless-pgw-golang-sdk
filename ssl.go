package ssl_wireless_pgw_golang_sdk

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/sagar290/ssl_wireless_pgw_golang_sdk/helpers"
	"github.com/sagar290/ssl_wireless_pgw_golang_sdk/structs"
	"log"
	"net/http"
)

type Ssl struct {
	StoreID                            string
	StorePassword                      string
	BaseUrl                            string
	InitURL                            string
	ValidationURL                      string
	RefundURL                          string
	RefundQueryURL                     string
	TransactionQueryBySessionIdURL     string
	TransactionQueryByTransactionIdURL string
}

func Init(storeId string, storePassword string, isLive bool) *Ssl {

	var url string

	if isLive {
		url = "https://securepay.sslcommerz.com"
	} else {
		url = "https://sandbox.sslcommerz.com"
	}

	return &Ssl{
		StoreID:                            storeId,
		StorePassword:                      storePassword,
		BaseUrl:                            url,
		InitURL:                            url + "/gwprocess/v4/api.php?",
		ValidationURL:                      url + "/validator/api/validationserverAPI.php?",
		RefundURL:                          url + "/validator/api/merchantTransIDvalidationAPI.php?",
		RefundQueryURL:                     url + "/validator/api/merchantTransIDvalidationAPI.php?",
		TransactionQueryBySessionIdURL:     url + "/validator/api/merchantTransIDvalidationAPI.php?",
		TransactionQueryByTransactionIdURL: url + "/validator/api/merchantTransIDvalidationAPI.php?",
	}
}

func (ssl *Ssl) MakePayment(postBody structs.PaymentBody) (structs.PaymentResponse, error) {

	var res structs.PaymentResponse

	// set store id and password
	postBody.StoreId = ssl.StoreID
	postBody.StorePasswd = ssl.StorePassword

	values := helpers.StructToMap(&postBody)

	response, err := http.PostForm(ssl.InitURL, values)

	if err != nil {
		log.Fatal(err)
		return structs.PaymentResponse{}, err
	}

	json.NewDecoder(response.Body).Decode(&res)
	// log.Printf("%v\n", res)
	return res, nil

}

func (ssl *Ssl) ValidatePayment(data structs.OrderValidateBody) (structs.OrderValidateResponse, error) {
	var res structs.OrderValidateResponse

	// set store id and password
	data.StoreId = ssl.StoreID
	data.StorePasswd = ssl.StorePassword

	v, _ := query.Values(data)

	//resp, err := http.Get(ssl.ValidationURL + "val_id=" + data.ValId + "&store_id=" + data.StoreId + "&store_passwd=" + data.StorePasswd)
	resp, err := http.Get(ssl.ValidationURL + helpers.Encode(v))

	if err != nil {
		log.Fatal(err)
		return structs.OrderValidateResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func (ssl *Ssl) InitiateRefund(data structs.InitiateRefundBody) (structs.InitiateRefundResponse, error) {
	var res structs.InitiateRefundResponse

	// set store id and password
	data.StoreId = ssl.StoreID
	data.StorePasswd = ssl.StorePassword

	v, _ := query.Values(data)

	resp, err := http.Get(ssl.RefundURL + helpers.Encode(v))

	if err != nil {
		log.Fatal(err)
		return structs.InitiateRefundResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func (ssl *Ssl) RefundQuery(data structs.RefundQueryURLBody) (structs.RefundQueryURLResponse, error) {
	var res structs.RefundQueryURLResponse

	// set store id and password
	data.StoreId = ssl.StoreID
	data.StorePasswd = ssl.StorePassword

	v, _ := query.Values(data)

	resp, err := http.Get(ssl.RefundQueryURL + helpers.Encode(v))

	if err != nil {
		log.Fatal(err)
		return structs.RefundQueryURLResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func (ssl *Ssl) TransactionQueryBySessionId(data structs.TransactionQueryBySessionKeyBody) (structs.TransactionQueryBySessionKeyResponse, error) {
	var res structs.TransactionQueryBySessionKeyResponse

	// set store id and password
	data.StoreId = ssl.StoreID
	data.StorePasswd = ssl.StorePassword

	v, _ := query.Values(data)

	resp, err := http.Get(ssl.TransactionQueryBySessionIdURL + helpers.Encode(v))

	if err != nil {
		log.Fatal(err)
		return structs.TransactionQueryBySessionKeyResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func (ssl *Ssl) TransactionQueryByTransactionId(data structs.TransactionQueryByTransactionIdBody) (structs.TransactionQueryByTransactionIdResponse, error) {
	var res structs.TransactionQueryByTransactionIdResponse

	// set store id and password
	data.StoreId = ssl.StoreID
	data.StorePasswd = ssl.StorePassword

	v, _ := query.Values(data)

	resp, err := http.Get(ssl.TransactionQueryBySessionIdURL + helpers.Encode(v))

	if err != nil {
		log.Fatal(err)
		return structs.TransactionQueryByTransactionIdResponse{}, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}
