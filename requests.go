package tinder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Auth authorises the client and initiates the connection.
func (tinder *Tinder) Auth() error {
	Headers := tinder.Headers
	Headers.Add("facebook_token", tinder.Facebook["facebook_token"])
	Headers.Add("facebook_ID", tinder.Facebook["facebook_ID"])
	response, err := http.PostForm(tinder.Host+"/auth", Headers)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var me Me
	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}
	tinder.Headers.Del("facebook_token")
	tinder.Headers.Del("facebook_ID")
	tinder.Token = me.Token
	tinder.Me = me

	return nil
}

//UpdatePreferences allows you to change your preferences for searching.
func (tinder *Tinder) UpdatePreferences(gender string, minAge int, maxAge int, maxDistance int, bio string) error {
	var genderInt int
	switch gender {
	case "male":
		genderInt = 0
	case "female":
		genderInt = 1
	default:
		genderInt = 1
	}
	//JSONStruct holds the prefernces to be marshalled to JSON.
	JSONStruct := &Profile{
		Gender:      genderInt,
		MinAge:      minAge,
		MaxAge:      maxAge,
		MaxDistance: maxDistance,
		Bio:         bio,
	}
	//JSONData is the encoded JSON of JSONStruct
	JSONData, err := json.Marshal(JSONStruct)
	if err != nil {
		return err
	}
	//JSONReader creates a reder to insert into the POST request
	JSONReader := bytes.NewReader(JSONData)

	req, err := http.NewRequest("POST", tinder.Host+"/profile", JSONReader)
	if err != nil {
		return err
	}
	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

//Ping takes longitude and latitude and will ping the location.
func (tinder *Tinder) Ping(lat float32, lon float32) error {
	GeoStruct := &Geo{
		Lat: lat,
		Lon: lon,
	}
	GeoData, err := json.Marshal(GeoStruct)
	if err != nil {
		return err
	}

	GeoReader := bytes.NewReader(GeoData)

	req, err := http.NewRequest("POST", tinder.Host+"/user/ping", GeoReader)
	if err != nil {
		return err
	}
	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var GeoResp GeoResponse
	err = json.Unmarshal(data, &GeoResp)
	if err != nil {
		return err
	}

	if len(GeoResp.Error) != 0 {
		return errors.New(GeoResp.Error)
	}

	return nil
}

//Report will report a userID
func (tinder *Tinder) Report(ID string, Cause string) error {
	var CauseID int
	switch Cause {
	case "spam":
		CauseID = 1
	case "offensive":
		CauseID = 2
	default:
		return errors.New("Cause can only be spam or offensive")
	}

	ReportStruct := &Report{
		Cause: CauseID,
	}

	ReportData, err := json.Marshal(ReportStruct)
	if err != nil {
		return err
	}

	ReportReader := bytes.NewReader(ReportData)

	req, err := http.NewRequest("POST", tinder.Host+"/report/"+fmt.Sprintf("%s", ID), ReportReader)
	if err != nil {
		return err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var ReportResp ReportResponse
	err = json.Unmarshal(data, &ReportResp)
	if err != nil {
		return err
	}

	if len(ReportResp.Error) != 0 {
		return errors.New(ReportResp.Error)
	}

	return nil
}

//GetUpdates finds the latest updates such as new matches and new messages.
func (tinder *Tinder) GetUpdates() (UpdatesResponse, error) {
	var UpdatesEmpty UpdatesResponse
	UpdatesStruct := &Updates{
		Limit: 40,
	}
	UpdatesData, err := json.Marshal(UpdatesStruct)
	if err != nil {
		return UpdatesEmpty, err
	}

	UpdatesReader := bytes.NewReader(UpdatesData)

	req, err := http.NewRequest("POST", tinder.Host+"/updates", UpdatesReader)
	if err != nil {
		return UpdatesEmpty, err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return UpdatesEmpty, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return UpdatesEmpty, err
	}

	var UpdatesResp UpdatesResponse
	err = json.Unmarshal([]byte(data), &UpdatesResp)
	if err != nil {
		return UpdatesEmpty, err
	}

	if len(UpdatesResp.Error) != 0 {
		return UpdatesEmpty, errors.New(UpdatesResp.Error)
	}

	return UpdatesResp, nil
}

func swipe(tinder *Tinder, recID string, method string) (swipeResp SwipeResponse, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", tinder.Host, method, recID), nil)
	if err != nil {
		return swipeResp, err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return swipeResp, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return swipeResp, err
	}

	if err := json.Unmarshal([]byte(data), &swipeResp); err != nil {
		return swipeResp, err
	}

	switch v := swipeResp.MatchInternal.(type) {
	case bool:
		swipeResp.Match = false
	case map[string]interface{}:
		swipeResp.Match = true
		swipeResp.MatchDetails = v
	}

	return swipeResp, nil
}

//Like will 'swipe right' on the given ID
func (tinder *Tinder) Like(recID string) (match bool, err error) {
	swipeResp, err := swipe(tinder, recID, "like")
	if err != nil {
		return false, err
	}
	return swipeResp.Match, nil
}

//Pass wil 'swipe left' on the given ID
func (tinder *Tinder) Pass(recID string) error {
	_, err := swipe(tinder, recID, "like")
	if err != nil {
		return err
	}
	return nil
}

//GetRecommendations will get a list of people to like or pass on.
func (tinder *Tinder) GetRecommendations() (resp RecommendationsResponse, err error) {
	req, err := http.NewRequest("GET", tinder.Host+"/user/recs", nil)
	if err != nil {
		return resp, err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal([]byte(data), &resp); err != nil {
		return resp, err
	}

	if resp.Message == "recs timeout" {
		return resp, RecsTimeout

	} else if resp.Message == "recs exhausted" {
		return resp, RecsExhausted
	}

	return resp, nil
}

//SendMessage will send a message to the the given ID.
func (tinder *Tinder) SendMessage(userID string, message string) error {
	type empStruct struct {
		MatchID string `json:"match_ID"`
		Message string `json:"message"`
	}
	RecStruct := &empStruct{
		MatchID: userID,
		Message: message,
	}
	RecData, err := json.Marshal(RecStruct)
	if err != nil {
		return err
	}

	RecReader := bytes.NewReader(RecData)
	req, err := http.NewRequest("POST", tinder.Host+"/user/matches/"+userID, RecReader)
	if err != nil {
		return err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	resp := make(map[string]interface{})

	if err = json.Unmarshal([]byte(data), &resp); err != nil {
		return err
	}

	return nil
}

//SetRequiredHeaders ensures correct HTTP headers are set.
func (tinder *Tinder) SetRequiredHeaders(request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("User-Agent", "Tinder/3.0.4 (iPhone; iOS 7.1; Scale/2.00)")
	request.Header.Set("X-Auth-Token", tinder.Token)

	return request
}
