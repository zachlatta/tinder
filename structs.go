package tinder

import (
	"net/http"
	"net/url"
	"time"
)

//Tinder holds all vital information for interfacing with the API.
type Tinder struct {
	Host     string
	Facebook map[string]string
	Token    string
	Headers  url.Values
	Client   http.Client
	Me       Me
}

//Me contains information about your own bio.
type Me struct {
	Token string
	User  struct {
		ID           string `json:"_ID"`
		APIToken     string `json:"api_token"`
		Bio          string `json:"bio"`
		FullName     string `json:"full_name"`
		Name         string `json:"name"`
		Discoverable bool   `json:"discoverable"`
	} `json:"user"`
}

//Profile contains information about profiles.
type Profile struct {
	Gender      int    `json:"gender"`
	MinAge      int    `json:"age_filter_min"`
	MaxAge      int    `json:"age_filter_max"`
	MaxDistance int    `json:"distance_filter"`
	Bio         string `json:"bio,omitempty"`
}

//Geo holds georgraphical information.
type Geo struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

//GeoResponse holds the response for georgraphical pings.
type GeoResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

//Report holds the report cause.
type Report struct {
	Cause int `json:"cause"`
}

//ReportResponse holds the data of a response.
type ReportResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
}

//Updates holds the update limit
type Updates struct {
	Limit int `json:"limit"`
}

//UpdatesResponse holds the response for Updates.
type UpdatesResponse struct {
	Status  string `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	Matches []struct {
		ID                string `json:"_ID"`
		CommonFriendCount int    `json:"common_friend_count"`
		CommonLikeCount   int    `json:"common_like_count"`
		MessageCount      int    `json:"message_count"`
		Messages          []struct {
			ID        string `json:"_ID"`
			MatchID   string `json:"match_ID"`
			To        string `json:"to"`
			From      string `json:"from"`
			Message   string `json:"message,omitempty"`
			Sent      string `json:"sent_date"`
			Timestamp int64  `json:"timestamp"`
		} `json:"messages"`
		Person struct {
			ID       string `json:"_ID"`
			Bio      string `json:"bio"`
			Birth    string `json:"birth_date"`
			Gender   int    `json:"gender"`
			Name     string `json:"name"`
			PingTime string `json:"ping_time"`
		} `json:"person"`
	} `json:"matches"`
}

//SwipeResponse is the response to swiping, whether you matched, and if so the match details.
type SwipeResponse struct {
	Match         bool
	MatchDetails  map[string]interface{}
	MatchInternal interface{} `json:"match"` // Used to unmarshal
}

//ProcessedFile holds the data of a processed image
type ProcessedFile struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

//Photo holds information about photos
type Photo struct {
	ID               string          `json:"photos"`
	Main             interface{}     `json:"main"`
	Crop             string          `json:"crop"`
	FileName         string          `json:"fileName"`
	Extension        string          `json:"extension"`
	YDistancePercent float64         `json:"ydistance_percent"`
	XDistancePercent float64         `json:"xdistance_percent"`
	YOffsetPercent   float64         `json:"yoffset_percent"`
	XOffsetPercent   float64         `json:"xoffset_percent"`
	ProcessedFiles   []ProcessedFile `json:"processedFiles"`
}

//Recommendation holds the info of a Recommendation.
type Recommendation struct {
	ID                string    `json:"_ID"`
	Bio               string    `json:"bio"`
	Birth             time.Time `json:"birth_date"`
	BirthInfo         string    `json:"birth_date_info"`
	Gender            int       `json:"gender"`
	Name              string    `json:"name"`
	DistanceInMiles   int       `json:"distance_mi"`
	CommonLikeCount   int       `json:"common_like_count"`
	CommonFriendCount int       `json:"common_friend_count"`
	PingTime          string    `json:"ping_time"`
	Photos            []Photo   `json:"photos"`
}

//RecommendationsResponse is the response you get from Recommendation.
type RecommendationsResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Results []Recommendation `json:"results"`
}
