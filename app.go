/*
*
https://github.com/openatx/uiautomator2#app-management
*/
package uiautomator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
Install an app
TODO: api "/install" not work
*/
func (ua *UIAutomator) AppInstall(u string) func() (bool, error) {
	wrapError := func(err error) func() (bool, error) {
		return func() (bool, error) {
			return false, err
		}
	}
	requestURL := fmt.Sprintf("http://%s:%d/install", ua.config.Host, ua.config.Port)
	response, err := http.PostForm(
		requestURL,
		url.Values{
			"url": {u},
		},
	)
	if err != nil {
		return wrapError(err)
	}
	if response.StatusCode != http.StatusOK {
		err = boom(response)
		return wrapError(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return wrapError(err)
	}
	id := string(body)
	return func() (bool, error) {
		var InstallReturned struct {
			ID        string `json:"id"`
			TitalSize string `json:"titalSize"`
			CopySize  string `json:"copiedSize"`
			Messge    string `json:"message"`
		}
		requestURL = fmt.Sprintf("http://%s:%d/install/%s", ua.config.Host, ua.config.Port, id)
		response, err = http.Get(requestURL)
		if err != nil {
			return false, err
		}
		if response.StatusCode != http.StatusOK {
			err = boom(response)
			return false, err
		}
		err = json.NewDecoder(response.Body).Decode(&InstallReturned)
		if err != nil {
			return false, err
		}
		fmt.Printf("%v\n", InstallReturned)
		return true, nil
	}
}

/*
Launch app
*/
func (ua *UIAutomator) AppStart(packageName string) error {
	_, err := ua.Shell(
		[]string{
			"monkey", "-p", packageName, "-c",
			"android.intent.category.LAUNCHER", "1",
		},
		10,
	)

	return err
}

/*
Stop app
*/
func (ua *UIAutomator) AppStop(packageName string) error {
	_, err := ua.Shell(
		[]string{
			"am", "force-stop", packageName,
		},
		10,
	)

	return err
}
