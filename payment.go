package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
)

func (c *Client) DirectDebitPayment(currency, amount, referenceNo, productCode, orderTitle string, mcc *string, expireTime *int64, paymentOptions *[]map[string]interface{}, urlParams *[]map[string]string) (*DirectDebitPaymentResponse, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	currentMcc := defaultMcc
	if mcc != nil {
		currentMcc = *mcc
	}

	requestBody := map[string]interface{}{
		"partnerReferenceNo": referenceNo,
		"merchantId":         c.Config.MerchantId,
		"amount": map[string]string{
			"currency": currency,
			"value":    amount,
		},
		"validUpTo": c.getExpireTime(expireTime),
		"additionalInfo": map[string]interface{}{
			"productCode": productCode,
			"mcc":         currentMcc,
			"envInfo": map[string]interface{}{
				"sourcePlatform": SourcePlatformIPG,
				"terminalType":   TerminalTypeSystem,
			},
			"order": map[string]string{
				"orderTitle": orderTitle,
			},
		},
	}

	if paymentOptions != nil && len(*paymentOptions) > 0 {
		requestBody["payOptionDetails"] = *paymentOptions
	}

	if urlParams != nil && len(*urlParams) > 0 {
		requestBody["urlParams"] = urlParams
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLDirectDebitPayment), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "DirectDebitPayment",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLDirectDebitPayment)

	var result DirectDebitPaymentResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "DirectDebitPayment",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "DirectDebitPayment",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) QuickPay(currency, amount, referenceNo, productCode, orderTitle string, mcc *string, expireTime *int64, paymentOptions *[]map[string]interface{}) (*QuickPayResponse, error) {
	if err := c.EnsureB2BAccessToken(); err != nil {
		return nil, err
	}

	accessToken := c.b2bAccessToken.AccessToken
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	currentMcc := defaultMcc
	if mcc != nil {
		currentMcc = *mcc
	}

	requestBody := map[string]interface{}{
		"title":              orderTitle,
		"partnerReferenceNo": referenceNo,
		"merchantId":         c.Config.MerchantId,
		"amount": map[string]string{
			"currency": currency,
			"value":    amount,
		},
		"validUpTo": c.getExpireTime(expireTime),
		"additionalInfo": map[string]interface{}{
			"productCode": productCode,
			"mcc":         currentMcc,
			"envInfo": map[string]interface{}{
				"sourcePlatform": SourcePlatformIPG,
				"terminalType":   TerminalTypeSystem,
			},
		},
	}

	if paymentOptions != nil && len(*paymentOptions) > 0 {
		requestBody["payOptionDetails"] = *paymentOptions
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLQuickPay), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "QuickPay",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLQuickPay)

	var result QuickPayResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "QuickPay",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "QuickPay",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) CancelOrder(referenceNo string) (*CancelOrderRequest, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	requestBody := map[string]interface{}{
		"merchantId":                 c.Config.MerchantId,
		"originalPartnerReferenceNo": referenceNo,
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLCancelPayment), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "CancelPayment",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLCancelPayment)

	var result CancelOrderRequest

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "CancelPayment",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "CancelPayment",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) QueryPayment(referenceNo string) (*QueryPaymentResponse, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	requestBody := map[string]interface{}{
		"serviceCode":                "00",
		"merchantId":                 c.Config.MerchantId,
		"originalPartnerReferenceNo": referenceNo,
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLQueryPayment), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "QueryPayment",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLQueryPayment)

	var result QueryPaymentResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "QueryPayment",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "QueryPayment",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) GenerateQRIS(currency, amount, referenceNo string) (interface{}, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	requestBody := map[string]interface{}{
		"partnerReferenceNo": referenceNo,
		"merchantId":         c.Config.MerchantId,
		"amount": map[string]string{
			"currency": currency,
			"value":    amount,
		},
		"additionalInfo": map[string]interface{}{
			"terminalSource": "MER",
			"envInfo": map[string]interface{}{
				"sourcePlatform": SourcePlatformIPG,
				"terminalType":   TerminalTypeSystem,
			},
		},
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLGenerateQRIS), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "GenerateQRIS",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLGenerateQRIS)

	var result interface{}

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "GenerateQRIS",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "GenerateQRIS",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) FinishNotify(danaReferenceNo, referenceNo, amount, latestTransactionStatus, createdTime, finishedTime string) (interface{}, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	requestBody := map[string]interface{}{
		"merchantId":                 c.Config.MerchantId,
		"originalReferenceNo":        danaReferenceNo,
		"originalPartnerReferenceNo": referenceNo,
		"amount": map[string]string{
			"currency": "IDR",
			"value":    amount,
		},
		"latestTransactionStatus": latestTransactionStatus,
		"createdTime":             createdTime,
		"finishedTime":            finishedTime,
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLFinishNotify), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "FinishNotify",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    c.getChannelId(),
		"ORIGIN":        c.getOrigin(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLFinishNotify)

	var result interface{}

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "FinishNotify",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "FinishNotify",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) RefundOrder(orderId, refundId, currency, amount string) (*RefundOrderResponse, error) {
	timestamp := c.getTimestamp()
	requestId := c.getRequestId(nil)

	requestBody := map[string]interface{}{
		"merchantId":                 c.Config.MerchantId,
		"originalPartnerReferenceNo": orderId,
		"partnerRefundNo":            refundId,
		"refundAmount": map[string]string{
			"currency": currency,
			"value":    amount,
		},
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLRefund), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "RefundOrder",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": requestId,
		"X-SIGNATURE":   *signature,
		"CHANNEL-ID":    "95221",
		"ORIGIN":        c.getOrigin(),
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLRefund)

	var result RefundOrderResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "RefundOrder",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "RefundOrder",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) TransactionHistory(fromDateTime, toDateTime *string, customerAccessToken *AccessToken) (*string, *TransactionHistoryResponse, error) {
	timestamp := c.getTimestamp()
	externalId := c.getRequestId(nil)
	accessToken := c.getCustomerAccessToken(customerAccessToken)

	if accessToken == nil {
		return nil, nil, fmt.Errorf("customer access token is required")
	}

	requestBody := map[string]interface{}{
		"additionalInfo": map[string]interface{}{
			"types":       []string{"PAYMENT", "REFUND", "OFFLINE_TOPUP", "TOP_UP", "REBATE"},
			"statuses":    []string{"SUCCESS", "FAILED", "INIT", "PROCESSING", "CLOSED", "REVOKED"},
			"accessToken": accessToken.AccessToken,
		},
	}

	if fromDateTime != nil {
		requestBody["fromDateTime"] = *fromDateTime
	}

	if toDateTime != nil {
		requestBody["toDateTime"] = *toDateTime
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLTransactionList), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "GetCustomerTransactionList",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":           "application/json",
		"Authorization-Customer": fmt.Sprintf("%s %s", accessToken.TokenType, accessToken.AccessToken),
		"X-TIMESTAMP":            timestamp,
		"X-SIGNATURE":            *signature,
		"ORIGIN":                 c.getOrigin(),
		"X-PARTNER-ID":           c.Config.ClientId,
		"X-EXTERNAL-ID":          externalId,
		"X-IP-ADDRESS":           c.getIpAddress(),
		"X-DEVICE-ID":            c.getDeviceId(),
		"X-LATITUDE":             c.getLatitude(),
		"X-LONGITUDE":            c.getLongitude(),
		"CHANNEL-ID":             "95221",
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLTransactionList)

	var result TransactionHistoryResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "GetCustomerTransactionList",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "GetCustomerTransactionList",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &externalId, &result, nil
}
