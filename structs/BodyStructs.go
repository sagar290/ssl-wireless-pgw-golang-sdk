package structs

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type PaymentBody struct {

	// basic
	StoreId            string `json:"store_id"`
	StorePasswd        string `json:"store_passwd"`
	TotalAmount        float64 `json:"total_amount"`
	Currency           string `json:"currency" default:"BDT"`
	TranId             string `json:"tran_id"`
	ProductCategory    string `json:"product_category"`
	SuccessUrl         string `json:"success_url"` //done
	FailUrl            string `json:"fail_url"`
	CancelUrl          string `json:"cancel_url"`
	IpnUrl             string `json:"ipn_url"`

	// card details
	MultiCardName      string `json:"multi_card_name"`
	AllowedBin         string `json:"allowed_bin"`
	EmiOption          string `json:"emi_option"`
	EmiMaxInstOption   string `json:"emi_max_inst_option"`
	EmiSelectedInst    string `json:"emi_selected_inst"`
	EmiAllowOnly       string `json:"emi_allow_only"`

	// customer details
	CusName            string `json:"cus_name"`
	CusEmail           string `json:"cus_email"`
	CusAdd1            string `json:"cus_add_1"`
	CusAdd2            string `json:"cus_add_2"`
	CusCity            string `json:"cus_city"`
	CusState           string `json:"cus_state"`
	CusPostcode        string `json:"cus_postcode"`
	CusCountry         string `json:"cus_country"`
	CusPhone           string `json:"cus_phone"`
	CusFax             string `json:"cus_fax"`

	// Shipping
	ShippingMethod     string `json:"shipping_method"`
	NumOfItem          int `json:"num_of_item"`
	ShipName           string `json:"ship_name"`
	ShipAdd1           string `json:"ship_add_1"`
	ShipAdd2           string `json:"ship_add_2"`
	ShipCity           string `json:"ship_city"`
	ShipState          string `json:"ship_state"`
	ShipPostcode       string `json:"ship_postcode"`
	ShipCountry        string `json:"ship_country"`

	// product details
	ProductName        string `json:"product_name"`
	ProductProfile     string `json:"product_profile"`
	HoursTillDeparture string `json:"hours_till_departure"`
	FlightType         string `json:"flight_type"`
	Pnr                string `json:"pnr"`
	JourneyFromTo      string `json:"journey_from_to"`
	ThirdPartyBooking  string `json:"third_party_booking"`
	HotelName          string `json:"hotel_name"`
	LengthOfStay       string `json:"length_of_stay"`
	CheckInTime        string `json:"check_in_time"`
	HotelCity          string `json:"hotel_city"`
	ProductType        string `json:"product_type"`
	TopupNumber        string `json:"topup_number"`
	CountryTopup       string `json:"country_topup"`
	Cart               string `json:"cart"`
	ProductAmount      string `json:"product_amount"`
	Vat                string `json:"vat"`
	DiscountAmount     string `json:"discount_amount"`
	ConvenienceFee     string `json:"convenience_fee"`

	// others
	ValueA             string `json:"value_a"`
	ValueB             string `json:"value_b"`
	ValueC             string `json:"value_c"`
	ValueD             string `json:"value_d"`
}

type OrderValidateBody struct {
	ValId         string `json:"val_id"`
	StoreId       string `json:"store_id"`
	StorePasswd string `json:"store_passwd"`
	Format        string `json:"format"`
	V             string `json:"v"`
}

type TransactionQueryBySessionKeyBody struct {
	Sessionkey  string `json:"sessionkey"`
	StoreId     string `json:"store_id"`
	StorePasswd string `json:"store_passwd"`
}

type TransactionQueryByTransactionIdBody struct {
	Sessionkey  string `json:"sessionkey"`
	StoreId     string `json:"store_id"`
	StorePasswd string `json:"store_passwd"`
}

type InitiateRefundBody struct {
	BankTranId string `json:"bank_tran_id"`
	StoreId       string `json:"store_id"`
	StorePasswd string `json:"store_passwd"`
	RefundAmount string `json:"refund_amount"`
	RefeId string `json:"refe_id"`
	Format string `json:"format"`
}

type RefundQueryURLBody struct {
	RefundRefId string `json:"refund_ref_id"`
	StoreId       string `json:"store_id"`
	StorePasswd string `json:"store_passwd"`
}

type CardDetails struct {
	MultiCardName      string
	AllowedBin         string
	EmiOption          string
	EmiMaxInstOption   string
	EmiSelectedInst    string
	EmiAllowOnly       string
}

func (ssl *PaymentBody) SetCardDetails(value CardDetails)  {
	ssl.MultiCardName      = value.MultiCardName
	ssl.AllowedBin         = value.AllowedBin
	ssl.EmiOption          = value.EmiOption
	ssl.EmiMaxInstOption   = value.EmiMaxInstOption
	ssl.EmiSelectedInst    = value.EmiSelectedInst
	ssl.EmiAllowOnly       = value.EmiAllowOnly
}

type ProductDetails struct {
	ProductName        string
	ProductProfile     string
	HoursTillDeparture string
	FlightType         string
	Pnr                string
	JourneyFromTo      string
	ThirdPartyBooking  string
	HotelName          string
	LengthOfStay       string
	CheckInTime        string
	HotelCity          string
	ProductType        string
	TopupNumber        string
	CountryTopup       string
	Cart               string
	ProductAmount      string
	Vat                string
	DiscountAmount     string
	ConvenienceFee     string
}

func (ssl *PaymentBody) SetProductDetails(value ProductDetails)  {
	ssl.ProductName        = value.ProductName
	ssl.ProductProfile     = value.ProductProfile
	ssl.HoursTillDeparture = value.HoursTillDeparture
	ssl.FlightType         = value.FlightType
	ssl.Pnr                = value.Pnr
	ssl.JourneyFromTo      = value.JourneyFromTo
	ssl.ThirdPartyBooking  = value.ThirdPartyBooking
	ssl.HotelName          = value.HotelName
	ssl.LengthOfStay       = value.LengthOfStay
	ssl.CheckInTime        = value.CheckInTime
	ssl.HotelCity          = value.HotelCity
	ssl.ProductType        = value.ProductType
	ssl.TopupNumber        = value.TopupNumber
	ssl.CountryTopup       = value.CountryTopup
	ssl.Cart               = value.Cart
	ssl.ProductAmount      = value.ProductAmount
	ssl.Vat                = value.Vat
	ssl.DiscountAmount     = value.DiscountAmount
	ssl.ConvenienceFee     = value.ConvenienceFee
}

type ShippingDetails struct {
	ShippingMethod     string
	NumOfItem          int
	ShipName           string
	ShipAdd1           string
	ShipAdd2           string
	ShipCity           string
	ShipState          string
	ShipPostcode       string
	ShipCountry        string
}

func (ssl *PaymentBody) SetShippingDetails(value ShippingDetails)  {
	ssl.ShippingMethod     = value.ShippingMethod
	ssl.NumOfItem          = value.NumOfItem
	ssl.ShipName           = value.ShipName
	ssl.ShipAdd1           = value.ShipAdd1
	ssl.ShipAdd2           = value.ShipAdd2
	ssl.ShipCity           = value.ShipCity
	ssl.ShipState          = value.ShipState
	ssl.ShipPostcode       = value.ShipPostcode
	ssl.ShipCountry        = value.ShipCountry
}

func (ssl *PaymentBody) SetTransactionID(ID string) {
	ssl.TranId = ID
}

func (ssl *PaymentBody) SetTotalAmount(price float64) {
	ssl.TotalAmount = price
}

func (ssl *PaymentBody) SetProductCategory(category string) {
	ssl.ProductCategory = category
}

func (ssl *PaymentBody) SetCurrency(currency string) {
	ssl.Currency = currency
}

type CustomerDetails struct{
	Name            string
	Email           string
	Phone           string
	Add1            string
	Add2            string
	City            string
	State           string
	Postcode        string
	Country         string
	Fax             string
}

func (ssl *PaymentBody) SetCustomerDetails(value CustomerDetails)  {
	ssl.CusName = value.Name
	ssl.CusEmail = value.Email
	ssl.CusPhone = value.Phone
	ssl.CusAdd1 = value.Add1
	ssl.CusAdd2 = value.Add2
	ssl.CusCity = value.City
	ssl.CusState = value.State
	ssl.CusPostcode = value.Postcode
	ssl.CusCountry = value.Country
	ssl.CusFax = value.Fax
}


type Value struct {
	A string
	B string
	C string
	D string
}

func (ssl *PaymentBody) SetValue(value Value) {
	ssl.ValueA = value.A
	ssl.ValueB = value.B
	ssl.ValueC = value.C
	ssl.ValueD = value.D
}

func (ssl *PaymentBody) SetIpnUrl(url string) {
	ssl.IpnUrl = url
}

func (ssl *PaymentBody) SetSuccessUrl(url string)  {
	ssl.SuccessUrl = url
}

func (ssl *PaymentBody) SetCancelUrl(url string)  {
	ssl.CancelUrl = url
}

func (ssl *PaymentBody) SetFailUrl(url string) {
	ssl.FailUrl = url
}


func  (ssl *PaymentBody) MakePayment(url string) (PaymentResponse, error)  {

	var res PaymentResponse

	body, _ := json.Marshal(ssl)

	resp, err := http.Post(url + "/gwprocess/v4/api.php", "application/json",
		bytes.NewBuffer(body))

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}