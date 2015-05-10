package tinder

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"errors"
	"fmt"
)

func (tinder *Tinder) Auth() error {
	Headers := tinder.Headers
	Headers.Add("facebook_token", tinder.Facebook["facebook_token"])
	Headers.Add("facebook_id", tinder.Facebook["facebook_id"])
	response, err := http.PostForm(tinder.Host + "/auth", Headers)
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
	tinder.Headers.Del("facebook_id")
	tinder.Token = me.Token
	tinder.Me = me

	return nil
}

func (tinder *Tinder) UpdatePreferences(gender string, min_age int, max_age int, max_distance int, bio string) error {
	var gender_int int
	switch gender {
		case "male":
			gender_int = 0
		case "female":
			gender_int = 1
		default:
			gender_int = 1
	}

	JsonStruct := &Profile{
		Gender: gender_int,
		Min_Age: min_age,
		Max_Age: max_age,
		Max_Distance: max_distance,
		Bio: bio,
	}
	JsonData, err := json.Marshal(JsonStruct)
	if err != nil {
		return err
	}

	JsonReader := bytes.NewReader(JsonData)

	req, err := http.NewRequest("POST", tinder.Host + "/profile", JsonReader)
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

	req, err := http.NewRequest("POST", tinder.Host + "/user/ping", GeoReader)
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

func (tinder *Tinder) Report(Id string, Cause string) error {
	var Cause_Id int
	switch Cause {
		case "spam":
			Cause_Id = 1
		case "offensive":
			Cause_Id = 2
		default:
			return errors.New("Cause can only be spam or offensive")
	}

	ReportStruct := &Report{
		Cause: Cause_Id,
	}

	ReportData, err := json.Marshal(ReportStruct)
	if err != nil {
		return err
	}

	ReportReader := bytes.NewReader(ReportData)

	req, err := http.NewRequest("POST", tinder.Host + "/report/" + fmt.Sprintf("%s", Id), ReportReader)
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

	req, err := http.NewRequest("POST", tinder.Host + "/updates", UpdatesReader)
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
		fmt.Printf("\n%s\n\n", err)
		return UpdatesEmpty, err
	}

	if len(UpdatesResp.Error) != 0 {
		return UpdatesEmpty, errors.New(UpdatesResp.Error)
	}

	return UpdatesResp, nil
}

func (tinder *Tinder) Like(matchId string) (match bool, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/like/%s", tinder.Host, matchId), nil)
	if err != nil {
		return false, err
	}

	req = tinder.SetRequiredHeaders(req)
	response, err := tinder.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	var result SwipeResponse
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return false, err
	}

	return result.Match, nil
}

func (tinder *Tinder) SetRequiredHeaders(request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("User-Agent", "Tinder/3.0.4 (iPhone; iOS 7.1; Scale/2.00)")
	request.Header.Set("X-Auth-Token", tinder.Token)

	return request
}
