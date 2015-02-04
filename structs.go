package tinder

import (
	"net/http"
	"net/url"
)

type Tinder struct {
	Host      string
	Facebook  map[string]string
	Token     string
	Headers   url.Values
	Client    http.Client
	Me        Me
}

type Me struct {
	Token string
	User  struct {
		Id             string `json:"_id"`
		ApiToken       string `json:"api_token"`
		Bio            string `json:"bio"`
		FullName       string `json:"full_name"`
		Name           string `json:"name"`
		Discoverable   bool   `json:"discoverable"`
	} `json:"user"`
}

type Profile struct {
	Gender          int	`json:"gender"`
	Min_Age         int	`json:"age_filter_min"`
	Max_Age         int	`json:""age_filter_max`
	Max_Distance    int	`json:"distance_filter"`
	Bio             string	`json:"bio,omitempty"`
}

type Geo struct {
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

type GeoResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"`
}
