package tinder

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	Token	string
}

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
	tinder.Headers.Add("X-Auth-Token", user.Token)

	return nil
}
