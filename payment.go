package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
)

func (c *Client) DirectDebitPayment(requestId string, data *DirectDebitPaymentRequest) (interface{}, error) {
	data.MerchantId = c.Config.MerchantId
	timestamp := c.getTimestamp()
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLDirectDebitPayment)
	encodeRequestBody := c.encodeRequestData(data)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLDirectDebitPayment), encodeRequestBody, timestamp)

	signature, err := c.sign(strToSign)
	if err != nil {
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

	var result interface{}

	if _, err = goutil.SendHttpPost(requestUrl, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QuickPay(requestId string, data *QuickPayRequest) (*QuickPayResponse, error) {
	if err := c.EnsureAccessToken(); err != nil {
		return nil, err
	}

	accessToken := c.accessToken.AccessToken
	data.MerchantId = c.Config.MerchantId
	timestamp := c.getTimestamp()
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLQuickPay)
	encodeRequestBody := c.encodeRequestData(data)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLQuickPay), encodeRequestBody, timestamp)

	signature, err := c.sign(strToSign)
	if err != nil {
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

	var result QuickPayResponse

	if _, err = goutil.SendHttpPost(requestUrl, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CancelPayment(requestId string, data *CancelPaymentRequest) (interface{}, error) {
	data.MerchantId = c.Config.MerchantId
	timestamp := c.getTimestamp()
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLCancelPayment)
	encodeRequestBody := c.encodeRequestData(data)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLCancelPayment), encodeRequestBody, timestamp)

	signature, err := c.sign(strToSign)
	if err != nil {
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

	var result interface{}

	if _, err = goutil.SendHttpPost(requestUrl, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) QueryPayment(requestId string, data *QueryPaymentRequest) (*QueryPaymentResponse, error) {
	data.MerchantId = c.Config.MerchantId
	timestamp := c.getTimestamp()
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLQueryPayment)
	encodeRequestBody := c.encodeRequestData(data)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLQueryPayment), encodeRequestBody, timestamp)

	signature, err := c.sign(strToSign)
	if err != nil {
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

	var result QueryPaymentResponse

	if _, err = goutil.SendHttpPost(requestUrl, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
