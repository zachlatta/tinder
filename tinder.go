package tinder

import (
	"net/url"
)

type Tinder struct {
	Host      string
	Facebook  map[string]string
	Headers   url.Values
}

func Init(FacebookUserId string, FacebookToken string) *Tinder {
	Host := "https://api.gotinder.com"
	Facebook := make(map[string]string)
	Facebook["facebook_token"] = FacebookToken
	Facebook["facebook_id"] = FacebookUserId
	Values := &url.Values{
		"Content-Type": {"application/json"},
		"User-Agent": {"Tinder/3.0.4 (iPhone; iOS 7.1; Scale/2.00)"},
	}

	return &Tinder{
		Host: Host,
		Facebook: Facebook,
		Headers: *Values,
	}
}
