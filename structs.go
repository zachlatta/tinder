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
}

type User struct {
	Token string
}

type Profile struct {
	Gender		int	`json:"gender"`
	Min_Age		int	`json:"age_filter_min"`
	Max_Age		int	`json:""age_filter_max`
	Max_Distance	int	`json:"distance_filter"`
	Bio		string	`json:"bio,omitempty"`
}
