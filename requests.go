package tinder

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"errors"
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
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return err
	}
	tinder.Headers.Del("facebook_token")
	tinder.Headers.Del("facebook_id")
	tinder.Token = user.Token

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

func (tinder *Tinder) SetRequiredHeaders(request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("User-Agent", "Tinder/3.0.4 (iPhone; iOS 7.1; Scale/2.00)")
	request.Header.Set("X-Auth-Token", tinder.Token)

	return request
}
