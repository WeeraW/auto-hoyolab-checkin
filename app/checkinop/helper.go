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

	"github.com/brokiem/auto-hoyolab-checkin/app/configcheckin"
	"github.com/brokiem/auto-hoyolab-checkin/app/cookiereader"
	"github.com/brokiem/auto-hoyolab-checkin/app/myconsole"
	"github.com/zellyn/kooky"
)

func DoCheckIn(config configcheckin.CheckinConfig) (message string, err error) {

	if configcheckin.ConfigData.AutoHideWindow {
		myconsole.HideConsole()
	}

	claimResult, err := GetClaimedStatus(cookiereader.HoyolabToken[0], cookiereader.HoyolabLtuid[0], configcheckin.ConfigData.GenshinImpact.ActId)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if claimResult.Data.IsSign {
		return "you've already checked in today~", nil
	} else {
		_, err = ClaimReward(cookiereader.HoyolabToken[0], cookiereader.HoyolabLtuid[0], configcheckin.ConfigData.GenshinImpact.ActId)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
		return "You've claimed your reward today!", nil
	}
}

func GetClaimedStatus(token *kooky.Cookie, ltuid *kooky.Cookie, actId string) (ClaimResult, error) {
	fmt.Println("Checking claimed status...")
	var err error
	var result ClaimResult = ClaimResult{}

	req, _ := http.NewRequest("GET", "https://sg-hk4e-api.hoyolab.com/event/sol/info", nil)

	params := url.Values{}
	params.Add("act_id", actId)
	req.URL.RawQuery = params.Encode()

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+actId)
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
		fmt.Println(err.Error())
		return ClaimResult{}, err
	}

	fmt.Println(resp.Status)

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	err = json.Unmarshal([]byte(bodyString), &result)
	fmt.Println("--------------------------------")
	if err != nil {
		result.AppMessage = "Error: " + err.Error()
		return result, err
	}
	return result, nil
}

func ClaimReward(token *kooky.Cookie, ltuid *kooky.Cookie, actId string) (string, error) {
	var err error
	m := map[string]string{"act_id": actId}
	read, write := io.Pipe()
	go func() {
		json.NewEncoder(write).Encode(m)
		write.Close()
	}()

	req, _ := http.NewRequest("POST", "https://sg-hk4e-api.hoyolab.com/event/sol/sign", read)

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Origin", "https://act.hoyolab.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://act.hoyolab.com/ys/event/signin-sea-v3/index.html?act_id="+actId)

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
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println(resp.Status)
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var result map[string]interface{}
	json.Unmarshal([]byte(bodyString), &result)

	return reflect.ValueOf(result["message"]).String(), nil
}

// RandomSleepTime returns a random time.Duration between 0 and max seconds
func RandomSleepTime(max int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(max)) * time.Second
}
