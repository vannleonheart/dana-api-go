package dana

const (
	DefaultTimezone  = "Asia/Jakarta"
	TimestampFormat  = "2006-01-02T15:04:05+07:00"
	DefaultChannelId = "0"

	URLAccessToken        = "v1.0/access-token/b2b.htm"
	URLQuickPay           = "v1.0/quick-pay.htm"
	URLDirectDebitPayment = "v1.0/debit/payment.htm"
	URLQueryPayment       = "v1.0/debit/status.htm"
	URLCancelPayment      = "v1.0/debit/cancel.htm"

	CurrencyIDR = "IDR"

	UrlParamTypeNotification  = "NOTIFICATION"
	UrlParamTypePaymentReturn = "PAY_RETURN"

	PayMethodBalance               = "BALANCE"
	PayMethodCoupon                = "COUPON"
	PayMethodNetBanking            = "NET_BANKING"
	PayMethodCreditCard            = "CREDIT_CARD"
	PayMethodDebitCard             = "DEBIT_CARD"
	PayMethodVirtualAccount        = "VIRTUAL_ACCOUNT"
	PayMethodOTC                   = "OTC"
	PayMethodDirectDebitCreditCard = "DIRECT_DEBIT_CREDIT_CARD"
	PayMethodDirectDebitDebitCard  = "DIRECT_DEBIT_DEBIT_CARD"
	PayMethodOnlineCredit          = "ONLINE_CREDIT"
	PayMethodLoanCredit            = "LOAN_CREDIT"

	SourcePlatformIPG = "IPG"

	TerminalTypeApp    = "APP"
	TerminalTypeWeb    = "WEB"
	TerminalTypeWap    = "WAP"
	TerminalTypeSystem = "SYSTEM"

	VirtualAccountBNI     = "VIRTUAL_ACCOUNT_BNI"
	VirtualAccountBCA     = "VIRTUAL_ACCOUNT_BCA"
	VirtualAccountMandiri = "VIRTUAL_ACCOUNT_MANDIRI"
	VirtualAccountBRI     = "VIRTUAL_ACCOUNT_BRI"
	VirtualAccountBTPN    = "VIRTUAL_ACCOUNT_BTPN"
	VirtualAccountPanin   = "VIRTUAL_ACCOUNT_PANI"
	VirtualAccountCIMB    = "VIRTUAL_ACCOUNT_CIMB"
	VirtualAccountPermata = "VIRTUAL_ACCOUNT_BNLI"
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

type GeneralResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type AccessTokenResponse struct {
	GeneralResponse
	*AccessToken
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
}

type QuickPayRequest struct {
	PartnerReferenceNo string            `json:"partnerReferenceNo"`
	MerchantId         string            `json:"merchantId"`
	SubMerchantId      *string           `json:"subMerchantId,omitempty"`
	Amount             Money             `json:"amount"`
	ExternalStoreId    *string           `json:"externalStoreId,omitempty"`
	ValidUpTo          *string           `json:"validUpTo,omitempty"`
	Title              string            `json:"title"`
	UrlParams          *[]UrlParam       `json:"urlParams,omitempty"`
	PayOtionDetails    []PayOptionDetail `json:"payOptionDetails"`
	AdditionalInfo     *AdditionalInfo   `json:"additionalInfo,omitempty"`
}

type QuickPayResponse struct {
	GeneralResponse
	PartnerReferenceNo *string `json:"partnerReferenceNo"`
	ReferenceNo        *string `json:"referenceNo"`
	AdditionalInfo     *struct {
		VirtualAccountInfo struct {
			Signature                string `json:"signature"`
			VirtualAccountExpiryTime string `json:"virtualAccountExpiryTime"`
			VirtualAccountCode       string `json:"virtualAccountCode"`
		}
		ExtendInfo interface{} `json:"extendInfo"`
	} `json:"additionalInfo"`
}

type DirectDebitPaymentRequest struct {
	PartnerReferenceNo string            `json:"partnerReferenceNo"`
	MerchantId         string            `json:"merchantId"`
	SubMerchantId      *string           `json:"subMerchantId,omitempty"`
	Amount             Money             `json:"amount"`
	UrlParams          *[]UrlParam       `json:"urlParams,omitempty"`
	ExternalStoreId    *string           `json:"externalStoreId,omitempty"`
	ValidUpTo          *string           `json:"validUpTo,omitempty"`
	PointOfInitiation  *string           `json:"pointOfInitiation,omitempty"`
	DisabledPayMethods *string           `json:"disabledPayMethods,omitempty"`
	PayOtionDetails    []PayOptionDetail `json:"payOptionDetails"`
	AdditionalInfo     *AdditionalInfo   `json:"additionalInfo,omitempty"`
}

type DirectDebitPaymentResponse struct {
	GeneralResponse
	PartnerReferenceNo *string `json:"partnerReferenceNo"`
	ReferenceNo        *string `json:"referenceNo"`
	WebRedirectUrl     *string `json:"webRedirectUrl"`
}

type CancelPaymentRequest struct {
	MerchantId                 string  `json:"merchantId"`
	SubMerchantId              *string `json:"subMerchantId,omitempty"`
	OriginalPartnerReferenceNo string  `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	OriginalExternalId         *string `json:"originalExternalId,omitempty"`
	Reason                     *string `json:"reason,omitempty"`
	ExternalStoreId            *string `json:"externalStoreId,omitempty"`
	Amount                     *Money  `json:"amount,omitempty"`
}

type CancelPaymentResponse struct {
	GeneralResponse
	OriginalPartnerReferenceNo *string `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	CancelTime                 *string `json:"cancelTime,omitempty"`
}

type QueryPaymentRequest struct {
	MerchantId                 string  `json:"merchantId"`
	ServiceCode                string  `json:"serviceCode"`
	OriginalPartnerReferenceNo *string `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	OriginalExternalId         *string `json:"originalExternalId,omitempty"`
	TransactionDate            *string `json:"transactionDate,omitempty"`
	Amount                     *Money  `json:"amount,omitempty"`
	SubMerchantId              *string `json:"subMerchantId,omitempty"`
	ExternalStoreId            *string `json:"externalStoreId,omitempty"`
}

type QueryPaymentResponse struct {
	GeneralResponse
	OriginalPartnerReferenceNo *string `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	ServiceCode                *string `json:"serviceCode,omitempty"`
	TransAmount                *Money  `json:"transAmount,omitempty"`
	TransactionStatusDesc      *string `json:"transactionStatusDesc,omitempty"`
	Amount                     *Money  `json:"amount,omitempty"`
	LatestTransactionStatus    *string `json:"latestTransactionStatus,omitempty"`
	Title                      *string `json:"title,omitempty"`
}

type Money struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type UrlParam struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	IsDeeplink string `json:"isDeeplink"`
}

type AdditionalInfo struct {
	ProductCode string  `json:"productCode"`
	Order       Order   `json:"order,omitempty"`
	Mcc         string  `json:"mcc"`
	EnvInfo     EnvInfo `json:"envInfo"`
	ExtendInfo  *string `json:"extendInfo,omitempty"`
}

type Order struct {
	MerchantTransType *string       `json:"merchantTransType,omitempty"`
	Goods             *[]Goods      `json:"goods,omitempty"`
	ShippingInfo      *ShippingInfo `json:"shippingInfo,omitempty"`
	CreatedTime       *string       `json:"createdTime,omitempty"`
	ExtendInfo        *string       `json:"extendInfo,omitempty"`
	OrderTitle        *string       `json:"orderTitle,omitempty"`
	OrderMemo         *string       `json:"orderMemo,omitempty"`
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
	OrderTerminalType  string  `json:"orderTerminalType,omitempty"`
	TerminalType       string  `json:"terminalType"`
	OrderOsType        *string `json:"orderOsType,omitempty"`
	MerchantAppVersion *string `json:"merchantAppVersion,omitempty"`
	ExtendInfo         *string `json:"extendInfo,omitempty"`
}

type PayOptionDetail struct {
	PayMethod      string                         `json:"payMethod"`
	PayOption      string                         `json:"payOption"`
	TransAmount    Money                          `json:"transAmount"`
	FeeAmount      *Money                         `json:"feeAmount,omitempty"`
	CardToken      *string                        `json:"cardToken,omitempty"`
	AdditionalInfo *PayOptionDetailAdditionalInfo `json:"additionalInfo,omitempty"`
}

type PayOptionDetailAdditionalInfo struct {
	VirtualAccountExpiryTime string  `json:"virtualAccountExpiryTime"`
	VirtualAccountCode       *string `json:"virtualAccountCode,omitempty"`
}
