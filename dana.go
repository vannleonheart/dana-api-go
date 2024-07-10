package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"time"
)

func New(config Config) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) SetB2BAccessToken(accessToken *AccessToken) {
	c.b2bAccessToken = accessToken
}

func (c *Client) ClearB2BAccessToken() {
	c.b2bAccessToken = nil
}

func (c *Client) WithB2BAccessToken(accessToken *AccessToken) *Client {
	c.SetB2BAccessToken(accessToken)

	return c
}

func (c *Client) SetCustomerAccessToken(accessToken *AccessToken) {
	c.customerAccessToken = accessToken
}

func (c *Client) ClearCustomerAccessToken() {
	c.customerAccessToken = nil
}

func (c *Client) WithCustomerAccessToken(accessToken *AccessToken) *Client {
	c.SetCustomerAccessToken(accessToken)

	return c
}

func (c *Client) SetDeviceId(deviceId string) {
	c.deviceId = &deviceId
}

func (c *Client) ClearDeviceId() {
	c.deviceId = nil
}

func (c *Client) WithDeviceId(deviceId string) *Client {
	c.SetDeviceId(deviceId)

	return c
}

func (c *Client) SetOrigin(origin string) {
	c.origin = &origin
}

func (c *Client) ClearOrigin() {
	c.origin = nil
}

func (c *Client) WithOrigin(origin string) *Client {
	c.SetOrigin(origin)

	return c
}

func (c *Client) SetIpAddress(ipAddress string) {
	c.ipAddress = &ipAddress
}

func (c *Client) ClearIpAddress() {
	c.ipAddress = nil
}

func (c *Client) WithIpAddress(ipAddress string) *Client {
	c.SetIpAddress(ipAddress)

	return c
}

func (c *Client) SetLatitude(lat string) {
	c.lat = &lat
}

func (c *Client) ClearLatitude() {
	c.lat = nil
}

func (c *Client) WithLatitude(lat string) *Client {
	c.SetLatitude(lat)

	return c
}

func (c *Client) SetLongitude(lon string) {
	c.lon = &lon
}

func (c *Client) ClearLongitude() {
	c.lon = nil
}

func (c *Client) WithLongitude(lon string) *Client {
	c.SetLongitude(lon)

	return c
}

func (c *Client) SetRequestId(requestId string) {
	c.requestId = &requestId
}

func (c *Client) ClearRequestId() {
	c.requestId = nil
}

func (c *Client) WithRequestId(requestId string) *Client {
	c.SetRequestId(requestId)

	return c
}

func (c *Client) SetGeneratedRequestId() {
	requestId := fmt.Sprintf("%s%s", time.Now().Format("20060102"), GenerateRequestId(8, 16, goutil.NumCharset))

	c.SetRequestId(requestId)
}

func (c *Client) WithGeneratedRequestId() *Client {
	c.SetGeneratedRequestId()

	return c
}
