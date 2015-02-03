package tinder

import (
	"net/http"
	"net/url"
)

func Init(FacebookUserId string, FacebookToken string) *Tinder {
	Host := "https://api.gotinder.com"
	Facebook := make(map[string]string)
	Facebook["facebook_token"] = FacebookToken
	Facebook["facebook_id"] = FacebookUserId
	Values := &url.Values{}
	Client := &http.Client{}

	return &Tinder{
		Host: Host,
		Facebook: Facebook,
		Headers: *Values,
		Client: *Client,
	}
}
