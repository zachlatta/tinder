package tinder

import (
	"net/http"
	"net/url"
	"time"
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

type Report struct {
	Cause  int `json:"cause"`
}

type ReportResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"`
}

type Updates struct {
	Limit  int `json:"limit"`
}

type UpdatesResponse struct {
	Status  string `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	Matches  []struct {
		ID                 string `json:"_id"`
		CommonFriendCount  int    `json:"common_friend_count"`
		CommonLikeCount    int    `json:"common_like_count"`
		MessageCount       int    `json:"message_count"`
		Messages []struct {
			ID         string `json:"_id"`
			MatchID    string `json:"match_id"`
			To         string `json:"to"`
			From       string `json:"from"`
			Message    string `json:"message,omitempty"`
			Sent       string `json:"sent_date"`
			Timestamp  int64  `json:"timestamp"`
		} `json:"messages"`
		Person struct {
			ID        string `json:"_id"`
			Bio       string `json:"bio"`
			Birth     string `json:"birth_date"`
			Gender    int    `json:"gender"`
			Name      string `json:"name"`
			PingTime  string `json:"ping_time`
		} `json:"person"`
	} `json:"matches"`
}

type SwipeResponse struct {
	Match bool
	MatchDetails map[string]interface{}
	MatchInternal interface{} `json:"match"` // Used to unmarshal
}

type ProcessedFile struct {
	Width int `json:"width"`
	Height int `json:"height"`
	Url string `json:"url"`
}

type Photo struct {
	ID string `json:"photos"`
	Main interface{} `json:"main"`
	Crop string `json:"crop"`
	FileName string `json:"fileName"`
	Extension string `json:"extension"`
	YDistancePercent float64 `json:"ydistance_percent"`
	XDistancePercent float64 `json:"xdistance_percent"`
	YOffsetPercent float64 `json:"yoffset_percent"`
	XOffsetPercent float64 `json:"xoffset_percent"`
	ProcessedFiles []ProcessedFile `json:"processedFiles"`
}

type Recommendation struct {
	ID string `json:"_id"`
	Bio string `json:"bio"`
	Birth time.Time `json:"birth_date"`
	BirthInfo string `json:"birth_date_info"`
	Gender int `json:"gender"`
	Name string `json:"name"`
	DistanceInMiles int `json:"distance_mi"`
	CommonLikeCount int `json:"common_like_count"`
	CommonFriendCount int `json:"common_friend_count"`
	PingTime string `json:"ping_time"`
	Photos []Photo `json:"photos"`
}

type RecommendationsResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Results []Recommendation `json:"results"`
}
