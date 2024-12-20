package checkinop

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsole"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/browserutils/kooky"
)

func DoCheckIn(cookie cookiereader.CheckInCookieV2, config configcheckin.CheckinConfig) (message string, err error) {

	if configcheckin.ConfigData.AutoHideWindow {
		myconsole.HideConsole()
	}

	claimResult, err := GetClaimedStatusV2(cookie.Cookies, config)
	if err != nil {
		servicelogger.Error(err.Error())
		return "", err
	}
	if claimResult.Data.IsSign {
		return "you've already checked in today~", nil
	} else {
		_, err = ClaimRewardV2(cookie.Cookies, config)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		return "You've claimed your reward today!", nil
	}
}

func GetClaimedStatus(token *kooky.Cookie, ltuid *kooky.Cookie, config configcheckin.CheckinConfig) (ClaimResult, error) {
	servicelogger.Infof("[%s] Checking claimed status...", config.GameName)
	var err error
	var result ClaimResult = ClaimResult{}

	req, _ := http.NewRequest("GET", config.InfoUrl, nil)

	params := url.Values{}
	params.Add("act_id", config.ActId)
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+config.ActId)
	req.Header.Add("Cache-Control", "'max-age=0")

	req.AddCookie(&http.Cookie{
		Name:   token.Name,
		Value:  token.Value,
		MaxAge: token.MaxAge,
	})
	req.AddCookie(&http.Cookie{
		Name:   ltuid.Name,
		Value:  ltuid.Value,
		MaxAge: ltuid.MaxAge,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		servicelogger.Error(err.Error())
		return ClaimResult{}, err
	}

	servicelogger.Debug(resp.Status)

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	servicelogger.Debug(bodyString)

	err = json.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		result.AppMessage = "Error: " + err.Error()
		return result, err
	}
	if result.Retcode != 0 {
		result.AppMessage = "Error: " + result.Message
		return result, fmt.Errorf(result.Message)
	}
	return result, nil
}

func GetClaimedStatusV2(cookies []cookiereader.CookieData, config configcheckin.CheckinConfig) (ClaimResult, error) {
	servicelogger.Infof("[%s] Checking claimed status...", config.GameName)
	var err error
	var result ClaimResult = ClaimResult{}

	req, _ := http.NewRequest("GET", config.InfoUrl, nil)

	params := url.Values{}
	params.Add("act_id", config.ActId)
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+config.ActId)
	req.Header.Add("Cache-Control", "'max-age=0")

	// req.AddCookie(&http.Cookie{
	// 	Name:   token.Name,
	// 	Value:  token.Value,
	// 	MaxAge: token.MaxAge,
	// })
	// req.AddCookie(&http.Cookie{
	// 	Name:   ltuid.Name,
	// 	Value:  ltuid.Value,
	// 	MaxAge: ltuid.MaxAge,
	// })
	for _, cookie := range cookies {
		req.AddCookie(&http.Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			MaxAge: cookie.MaxAge,
		})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		servicelogger.Error(err.Error())
		return ClaimResult{}, err
	}

	servicelogger.Debug(resp.Status)

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	servicelogger.Debug(bodyString)

	err = json.Unmarshal([]byte(bodyString), &result)
	if err != nil {
		result.AppMessage = "Error: " + err.Error()
		return result, err
	}
	if result.Retcode != 0 {
		result.AppMessage = "Error: " + result.Message
		return result, fmt.Errorf(result.Message)
	}
	return result, nil
}

func ClaimReward(token *kooky.Cookie, ltuid *kooky.Cookie, config configcheckin.CheckinConfig) (string, error) {
	servicelogger.Infof("[%s] Claiming reward...", config.GameName)
	var err error
	m := map[string]string{"act_id": config.ActId}
	read, write := io.Pipe()
	go func() {
		json.NewEncoder(write).Encode(m)
		write.Close()
	}()

	req, _ := http.NewRequest("POST", config.SignUrl, read)

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+config.ActId)

	req.AddCookie(&http.Cookie{
		Name:   token.Name,
		Value:  token.Value,
		MaxAge: token.MaxAge,
	})
	req.AddCookie(&http.Cookie{
		Name:   ltuid.Name,
		Value:  ltuid.Value,
		MaxAge: ltuid.MaxAge,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		servicelogger.Error(err.Error())
		return "", err
	}
	servicelogger.Debug(resp.Status)
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	servicelogger.Debug(bodyString)
	var result map[string]interface{}
	json.Unmarshal([]byte(bodyString), &result)

	return reflect.ValueOf(result["message"]).String(), nil
}

func ClaimRewardV2(cookies []cookiereader.CookieData, config configcheckin.CheckinConfig) (ClaimResult, error) {
	servicelogger.Infof("[%s] Claiming reward...", config.GameName)
	var err error
	m := map[string]string{"act_id": config.ActId}
	read, write := io.Pipe()
	go func() {
		json.NewEncoder(write).Encode(m)
		write.Close()
	}()

	req, _ := http.NewRequest("POST", config.SignUrl, read)

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+config.ActId)

	for _, cookie := range cookies {
		req.AddCookie(&http.Cookie{
			Name:    cookie.Name,
			Value:   cookie.Value,
			MaxAge:  cookie.MaxAge,
			Expires: cookie.Expires,
			Domain:  cookie.Domain,
		})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		servicelogger.Error(err.Error())
		return ClaimResult{}, err
	}
	servicelogger.Debug(resp.Status)
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	servicelogger.Debug(bodyString)
	var result ClaimResult
	json.Unmarshal([]byte(bodyString), &result)

	return result, nil
}

// RandomSleepTime returns a random time.Duration between min and max seconds
func RandomSleepTime(min int, max int) time.Duration {
	return time.Duration(rand.Intn(max-min)+min) * time.Second
}
