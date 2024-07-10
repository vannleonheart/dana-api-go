package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
)

func (c *Client) DirectDebitPayment(currency, amount, referenceNo, productCode, orderTitle string, mcc *string, paymentOptions *[]map[string]interface{}, urlParams *[]map[string]string) (*DirectDebitPaymentResponse, error) {
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

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
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

func (c *Client) QuickPay(currency, amount, referenceNo, productCode, orderTitle string, mcc *string, paymentOptions *[]map[string]interface{}) (*QuickPayResponse, error) {
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

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
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

func (c *Client) CancelPayment(referenceNo string) (*CancelPaymentResponse, error) {
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

	var result CancelPaymentResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
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

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
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

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
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
