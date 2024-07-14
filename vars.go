package dana

const (
	defaultTimezone  = "Asia/Jakarta"
	TimestampFormat  = "2006-01-02T15:04:05+07:00"
	defaultDevideId  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
	defaultChannelId = "0"
	defaultMcc       = "412"

	URLAccessToken        = "v1.0/access-token/b2b.htm"
	URLQuickPay           = "v1.0/quick-pay.htm"
	URLDirectDebitPayment = "v1.0/debit/payment.htm"
	URLQueryPayment       = "v1.0/debit/status.htm"
	URLCancelPayment      = "v1.0/debit/cancel.htm"
	URLGenerateQRIS       = "v1.0/qr/qr-mpm-generate.htm"
	URLGetAuthCode        = "v1.0/get-auth-code"
	URLApplyToken         = "v1.0/access-token/b2b2c.htm"
	URLApplyOTT           = "v1.0/qr/apply-ott.htm"
	URLUnbindToken        = "v1.0/registration-account-unbinding.htm"
	URLBalanceInquiry     = "v1.0/balance-inquiry.htm"
	URLFinishNotify       = "v1.0/debit/notify"
	URLRefund             = "v1.0/debit/refund.htm"
	URLTransactionList    = "v1.0/transaction-history-list.htm"

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
	Config              Config
	b2bAccessToken      *AccessToken
	customerAccessToken *AccessToken
	origin              *string
	ipAddress           *string
	lat                 *string
	lon                 *string
	requestId           *string
	deviceId            *string
}

type Config struct {
	ApiUrl               string     `json:"api_url"`
	WebUrl               string     `json:"web_url"`
	MerchantId           string     `json:"merchant_id"`
	ClientId             string     `json:"client_id"`
	ClientSecret         string     `json:"client_secret"`
	PublicKey            string     `json:"public_key"`
	PrivateKey           string     `json:"private_key"`
	FinishPaymentUrl     string     `json:"finish_payment_url"`
	FinishRefundUrl      string     `json:"finish_refund_url"`
	FinishPaymentCodeUrl string     `json:"finish_payment_code_url"`
	FinishRedirectUrl    string     `json:"finish_redirect_url"`
	Timezone             string     `json:"timezone"`
	Origin               string     `json:"origin"`
	IpAddress            string     `json:"ip_address"`
	Latitude             string     `json:"latitude"`
	Longitude            string     `json:"longitude"`
	Log                  *LogConfig `json:"log,omitempty"`
}

type LogConfig struct {
	Enable    bool   `json:"enable"`
	Level     string `json:"level"`
	Path      string `json:"path"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
	Rotation  string `json:"rotation"`
}

type GeneralResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type GetB2BAccessTokenResponse struct {
	GeneralResponse
	*AccessToken
}

type AccessToken struct {
	AccessToken           string  `json:"accessToken"`
	TokenType             string  `json:"tokenType"`
	ExpiresIn             *int    `json:"expiresIn"`
	AccessTokenExpiryTime *string `json:"accessTokenExpiryTime"`
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

type DirectDebitPaymentResponse struct {
	GeneralResponse
	PartnerReferenceNo *string `json:"partnerReferenceNo"`
	ReferenceNo        *string `json:"referenceNo"`
	WebRedirectUrl     *string `json:"webRedirectUrl"`
}

type CancelPaymentResponse struct {
	GeneralResponse
	OriginalPartnerReferenceNo *string `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string `json:"originalReferenceNo,omitempty"`
	CancelTime                 *string `json:"cancelTime,omitempty"`
}

type QueryPaymentResponse struct {
	GeneralResponse
	OriginalPartnerReferenceNo *string                 `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        *string                 `json:"originalReferenceNo,omitempty"`
	OriginalExternalId         *string                 `json:"originalExternalId,omitempty"`
	ServiceCode                *string                 `json:"serviceCode,omitempty"`
	TransAmount                *Money                  `json:"transAmount,omitempty"`
	TransactionStatusDesc      *string                 `json:"transactionStatusDesc,omitempty"`
	Amount                     *Money                  `json:"amount,omitempty"`
	FeeAmount                  *Money                  `json:"feeAmount,omitempty"`
	LatestTransactionStatus    *string                 `json:"latestTransactionStatus,omitempty"`
	Title                      *string                 `json:"title,omitempty"`
	OriginalResponseCode       *string                 `json:"originalResponseCode,omitempty"`
	OriginalResponseMessage    *string                 `json:"originalResponseMessage,omitempty"`
	SessionId                  *string                 `json:"sessionId,omitempty"`
	RequestId                  *string                 `json:"requestId,omitempty"`
	PaidTime                   *string                 `json:"paidTime,omitempty"`
	AdditionalInfo             *map[string]interface{} `json:"additionalInfo,omitempty"`
}

type GenerateQRISRequest struct {
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      *string         `json:"subMerchantId,omitempty"`
	StoreId            *string         `json:"storeId,omitempty"`
	TerminalId         *string         `json:"terminalId,omitempty"`
	PartnerReferenceNo string          `json:"partnerReferenceNo "`
	Amount             Money           `json:"amount,omitempty"`
	FeeAmount          *Money          `json:"feeAmount,omitempty"`
	ValidityPeriod     *string         `json:"validityPeriod,omitempty"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo,omitempty"`
}

type CustomerApplyTokenResponse struct {
	GeneralResponse
	*AccessToken
	RefreshToken           string                 `json:"refreshToken"`
	RefreshTokenExpiryTime string                 `json:"refreshTokenExpiryTime"`
	AdditionalInfo         map[string]interface{} `json:"additionalInfo"`
}

type CustomerApplyOTTResponse struct {
	GeneralResponse
	ResourceType string `json:"resourceType"`
	Value        string `json:"value"`
}

type CancelOrderRequest struct {
	GeneralResponse
	OriginalReferenceNo        string `json:"originalReferenceNo"`
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
	CancelTime                 string `json:"cancelTime"`
}

type RefundOrderResponse struct {
	GeneralResponse
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo        string `json:"originalReferenceNo"`
	PartnerRefundNo            string `json:"partnerRefundNo"`
	RefundNo                   string `json:"refundNo"`
	RefundTime                 string `json:"refundTime"`
	RefundAmount               Money  `json:"refundAmount"`
}

type TransactionHistoryResponse struct {
	GeneralResponse
	DetailData []struct {
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		DateTime           string `json:"dateTime"`
		Amount             Money  `json:"amount"`
		Type               string `json:"type"`
		Status             string `json:"status"`
		SourceOfFunds      []struct {
			Source string `json:"source"`
			Amount Money  `json:"amount"`
		}
		AdditionalInfo map[string]interface{} `json:"additionalInfo"`
	} `json:"detailData"`
	AdditionalInfo *struct {
		Paginator struct {
			PageNum    string `json:"pageNum"`
			PageSize   string `json:"pageSize"`
			TotalCount string `json:"totalCount"`
			TotalPage  string `json:"totalPage"`
		}
	} `json:"additionalInfo"`
}

type Money struct {
	Currency      string  `json:"currency"`
	Value         string  `json:"value"`
	ExternalId    string  `json:"externalId"`
	ChannelId     string  `json:"channelId"`
	MerchantId    string  `json:"merchantId"`
	SubMerchantId *string `json:"subMerchantId"`
}

type UrlParam struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	IsDeeplink string `json:"isDeeplink"`
}

type AdditionalInfo struct {
	ProductCode    *string  `json:"productCode,omitempty"`
	Order          *Order   `json:"order,omitempty"`
	Mcc            *string  `json:"mcc,omitempty"`
	EnvInfo        *EnvInfo `json:"envInfo,omitempty"`
	ExtendInfo     *string  `json:"extendInfo,omitempty"`
	TerminalSource *string  `json:"terminalSource,omitempty"`
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
	SourcePlatform     *string `json:"sourcePlatform,omitempty"`
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
