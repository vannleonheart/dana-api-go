package dana

const (
	DefaultTimezone = "Asia/Jakarta"
	TimestampFormat = "2006-01-02T15:04:05+07:00"

	URLAccessToken     = "v1.0/access-token/b2b.htm"
	URLPaymentRedirect = "v1.0/payment-gateway/payment.htm"

	CurrencyIDR = "IDR"

	UrlTypeNotification  = "NOTIFICATION"
	UrlTypePaymentReturn = "PAY_RETURN"

	ScenarioRedirection = "REDIRECTION"
	ScenarioApi         = "API"
)

type Client struct {
	Config      Config
	accessToken *AccessToken
}

type Config struct {
	BaseUrl              string  `json:"base_url"`
	MerchantId           string  `json:"merchant_id"`
	ClientId             string  `json:"client_id"`
	ClientSecret         string  `json:"client_secret"`
	PublicKey            string  `json:"public_key"`
	PrivateKey           string  `json:"private_key"`
	FinishPaymentUrl     *string `json:"finish_payment_url"`
	FinishRefundUrl      *string `json:"finish_refund_url"`
	FinishPaymentCodeUrl *string `json:"finish_payment_code_url"`
	FinishRedirectUrl    *string `json:"finish_redirect_url"`
	Timezone             *string `json:"timezone"`
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
}

type AccessTokenResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	*AccessToken
}

type CreatePaymentRedirectRequest struct {
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      *string         `json:"subMerchantId,omitempty"`
	Amount             Money           `json:"amount"`
	ExternalStoreId    *string         `json:"externalStoreId,omitempty"`
	ValidUpTo          *string         `json:"validUpTo,omitempty"`
	DisabledPayMethods *string         `json:"disabledPayMethods,omitempty"`
	UrlParams          *[]UrlParams    `json:"urlparams,omitempty"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type CreatePaymentRedirectResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type UrlParams struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	IsDeeplink string `json:"IsDeeplink"`
}

type Money struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type AdditionalInfo struct {
	Order      Order   `json:"order,omitempty"`
	Mcc        string  `json:"mcc,omitempty"`
	ExtendInfo *string `json:"extendInfo,omitempty"`
	EnvInfo    EnvInfo `json:"envInfo,omitempty"`
}

type Order struct {
	MerchantTransType *string       `json:"merchantTransType,omitempty"`
	OrderTitle        string        `json:"orderTitle"`
	Scenario          string        `json:"scenario"`
	Goods             []Goods       `json:"goods"`
	Buyer             Buyer         `json:"buyer"`
	ShippingInfo      *ShippingInfo `json:"shippingInfo,omitempty"`
	ExtendInfo        *string       `json:"extendInfo,omitempty"`
}

type Goods struct {
	Unit               *string `json:"unit,omitempty"`
	Category           string  `json:"category"`
	Price              Money   `json:"price"`
	MerchantShippingId *string `json:"merchantShippingId,omitempty"`
	MerchantGoodsId    string  `json:"merchantGoodsId"`
	Description        string  `json:"description"`
	SnapshotUrl        *string `json:"snapshotUrl,omitempty"`
	Quantity           string  `json:"quantity"`
	ExtendInfo         *string `json:"extendInfo,omitempty"`
}

type Buyer struct {
	ExternalUserId   *string `json:"externalUserId,omitempty"`
	ExternalUserType *string `json:"externalUserType,omitempty"`
	UserId           *string `json:"userId,omitempty"`
	Nickname         *string `json:"nickname,omitempty"`
}

type ShippingInfo struct {
	ChargeAmount       *Money  `json:"chargeAmount"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	TrackingNo         *string `json:"trackingNo,omitempty"`
	CountryName        string  `json:"countryName"`
	MerchantShippingId string  `json:"merchantShippingId"`
	CityName           string  `json:"cityName"`
	Address1           string  `json:"address1"`
	Address2           *string `json:"address2,omitempty"`
	PhoneNo            *string `json:"phoneNo,omitempty"`
	AreaName           *string `json:"areaName,omitempty"`
	Email              *string `json:"email,omitempty"`
	ZipCode            string  `json:"zipCode"`
	StateName          string  `json:"stateName"`
	FaxNo              *string `json:"faxNo,omitempty"`
	Carrier            *string `json:"carrier,omitempty"`
	MobileNo           *string `json:"mobileNo,omitempty"`
}

type EnvInfo struct {
	SessionId          *string `json:"sessionId,omitempty"`
	TokenId            *string `json:"tokenId,omitempty"`
	WebsiteLanguage    *string `json:"websiteLanguage,omitempty"`
	ClientIp           *string `json:"clientIp,omitempty"`
	OSType             *string `json:"osType,omitempty"`
	AppVersion         *string `json:"appVersion,omitempty"`
	SDKVersion         *string `json:"sdkVersion,omitempty"`
	SourcePlatform     string  `json:"sourcePlatform"`
	ClientKey          *string `json:"clientKey,omitempty"`
	OrderTerminalType  string  `json:"orderTerminalType"`
	TerminalType       string  `json:"terminalType"`
	OrderOsType        *string `json:"orderOsType,omitempty"`
	MerchantAppVersion *string `json:"merchantAppVersion,omitempty"`
	ExtendInfo         *string `json:"extendInfo,omitempty"`
}
