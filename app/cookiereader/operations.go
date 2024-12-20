package cookiereader

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"slices"

	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/mynotify"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/browserutils/kooky"
)

// CheckInCookieV2 is a struct to store cookies key format <browser>:<profile>
var HoyolabCookies = map[string]CheckInCookieV2{}

// ReadCookie reads cookies from browser or file
func ReadCookie() error {
	err := ReadCookiesFromFile()
	if err != nil {
		fmt.Println(err)
		err = ReadCookieFromBrowser()
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadCookieFromBrowser() error {
	defer onReadFromBrowserPanic()

	servicelogger.Info("Reading cookies from browser...")
	mynotify.Notify("Reading cookies from browser...")
	cookies := kooky.TraverseCookies(context.TODO(), kooky.Valid, kooky.DomainHasSuffix(`.hoyolab.com`))
	isFoundCredential := false
	for cookie := range cookies {
		if cookie == nil {
			continue
		}
		servicelogger.Info("Reading cookies from " + cookie.Browser.Browser())
		keyHoyoLab := fmt.Sprintf("%s:%s", cookie.Browser.Browser(), cookie.Browser.Profile())
		hoyolabCookyStore, ok := HoyolabCookies[keyHoyoLab]
		if !ok {
			hoyolabCookyStore = CheckInCookieV2{}
		}
		if slices.Contains(myconsants.CookieNames, cookie.Name) {
			isFoundCredential = true
			hoyolabCookyStore.Cookies = append(hoyolabCookyStore.Cookies, toCookieData(*cookie))
		}
		if !isFoundCredential {
			isFoundCredential = len(hoyolabCookyStore.Cookies) > 0
		}
		hoyolabCookyStore.Expire = cookie.Expires.Unix()
		HoyolabCookies[keyHoyoLab] = hoyolabCookyStore
	}
	if !isFoundCredential {
		return fmt.Errorf("account credential not found, please login to hoyolab once in Chrome/Firefox/Opera/Safari or close browser and try again")
	}
	// write to file
	servicelogger.Info("Writing cookies to file...")
	mynotify.Notify("Writing cookies to file...")
	errWrite := WriteCookiesToFile()
	if errWrite != nil {
		return errWrite
	}
	servicelogger.Info("Cookies loaded!")
	mynotify.Notify("Cookies loaded!")
	return nil
}

func toCookieData(values kooky.Cookie) CookieData {
	return CookieData{
		Name:             values.Name,
		Value:            values.Value,
		Quoted:           values.Quoted,
		Path:             values.Path,
		Domain:           values.Domain,
		Expires:          values.Expires,
		MaxAge:           values.MaxAge,
		Secure:           values.Secure,
		HttpOnly:         values.HttpOnly,
		SameSite:         int(values.SameSite),
		Partitioned:      values.Partitioned,
		Creation:         values.Creation,
		Browser:          values.Browser.Browser(),
		Profile:          values.Browser.Profile(),
		IsDefaultProfile: values.Browser.IsDefaultProfile(),
		FilePath:         values.Browser.FilePath(),
	}
}

func filterHoloLabCookies(values []CheckInCookieV2) []CheckInCookieV2 {
	var result []CheckInCookieV2 = []CheckInCookieV2{}
	for _, cookie := range values {
		if len(cookie.Cookies) == len(myconsants.CookieNames) {
			result = append(result, cookie)
		}
	}
	return result
}

func ReadCookiesFromFile() error {
	defer onPanic()

	servicelogger.Info("Reading cookies from file...")
	mynotify.Notify("Reading cookies from file...")
	if _, err := os.Stat("cookieconfig.json"); err != nil {
		return fmt.Errorf("account credential not found, please login to hoyolab once in Chrome/Firefox/Opera/Safari")
	}
	jsonFile, _ := os.Open("cookieconfig.json")
	fileInfo, err := jsonFile.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() <= 0 {
		return fmt.Errorf("account credential not found, please login to hoyolab once in Chrome/Firefox/Opera/Safari")
	}
	byteValue, err := os.ReadFile("cookieconfig.json")
	if err != nil {
		return err
	}
	var result map[string]CheckInCookieV2 = map[string]CheckInCookieV2{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}
	HoyolabCookies = result
	servicelogger.Info("Cookies loaded!")
	mynotify.Notify("Cookies loaded!")
	return nil
}

func WriteCookiesToFile() error {
	defer onPanic()

	var err error
	file, err := os.OpenFile("cookieconfig.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	dataJSON, err := json.Marshal(HoyolabCookies)
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(dataJSON))
	if err != nil {
		return err
	}
	return nil
}

func onReadFromBrowserPanic() {
	if r := recover(); r != nil {
		servicelogger.Error(fmt.Sprintf("Panic: on read cookie from browser %v", r))
		servicelogger.Fatal(string(debug.Stack()))
	}
}

func onPanic() {
	if r := recover(); r != nil {
		servicelogger.Error(fmt.Sprintf("Panic: %v", r))
		servicelogger.Fatal(string(debug.Stack()))
	}
}
