package cookiereader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/gen2brain/beeep"
	"github.com/zellyn/kooky"
)

var HoyolabCookies = []CheckInCookie{}

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
	servicelogger.Info("Reading cookies from browser...")
	beeep.Notify(myconsants.AppName, "Reading cookies from browser...", "")
	cookieStores := kooky.FindAllCookieStores()
	isFoundCredential := false
	for _, cookieStore := range cookieStores {
		servicelogger.Info("Reading cookies from " + cookieStore.Browser())
		resultToken, errRead := cookieStore.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.hoyolab.com`), kooky.Name(`ltoken`))
		if errRead != nil {
			servicelogger.Error(errRead.Error())
			continue
		}
		if len(resultToken) <= 0 {
			servicelogger.Info("No cookies found in " + cookieStore.Browser())
			continue
		}

		resultLtuid, errRead := cookieStore.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.hoyolab.com`), kooky.Name(`ltuid`))
		if errRead != nil {
			servicelogger.Error(errRead.Error())
			continue
		}
		if len(resultLtuid) <= 0 {
			servicelogger.Info("No cookies found in " + cookieStore.Browser())
			continue
		}
		servicelogger.Info("Found " + fmt.Sprint(len(resultToken)) + " cookies")

		for index, cookie := range resultToken {
			HoyolabCookies = append(HoyolabCookies, CheckInCookie{
				Ltuid:  *resultLtuid[index],
				Token:  *cookie,
				Expire: cookie.Expires.Unix(),
			})
		}

		isFoundCredential = true
		break
	}
	if !isFoundCredential {
		return fmt.Errorf("account credential not found, please login to hoyolab once in Chrome/Firefox/Opera/Safari or close browser and try again")
	}
	// write to file
	servicelogger.Info("Writing cookies to file...")
	beeep.Notify(myconsants.AppName, "Writing cookies to file...", "")
	errWrite := WriteCookiesToFile()
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func ReadCookiesFromFile() error {
	servicelogger.Info("Reading cookies from file...")
	beeep.Notify(myconsants.AppName, "Reading cookies from file...", "")
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
	var result []CheckInCookie
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}
	HoyolabCookies = result
	servicelogger.Info("Cookies loaded!")
	beeep.Notify(myconsants.AppName, "Cookies loaded!", "")
	return nil
}

func WriteCookiesToFile() error {
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
