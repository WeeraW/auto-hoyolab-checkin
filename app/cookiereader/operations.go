package cookiereader

import (
	"fmt"

	"github.com/zellyn/kooky"
)

var HoyolabToken []*kooky.Cookie
var HoyolabLtuid []*kooky.Cookie
var CookiePath string = `C:\Users\weerapong\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies`

func ReadCookieFromBrowser() error {
	fmt.Println("Reading cookies from browser...")
	cookieStores := kooky.FindAllCookieStores()
	var token []*kooky.Cookie
	var ltuid []*kooky.Cookie
	isFoundCredential := false
	for _, cookieStore := range cookieStores {
		fmt.Println("Reading cookies from", cookieStore.Browser())
		resultToken, errRead := cookieStore.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.hoyolab.com`), kooky.Name(`ltoken`))
		if errRead != nil {
			fmt.Println("Error reading cookies from", cookieStore.Browser(), errRead)
			continue
		}
		if len(resultToken) <= 0 {
			fmt.Println("No cookies found in", cookieStore.Browser())
			continue
		}

		resultLtuid, errRead := cookieStore.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`.hoyolab.com`), kooky.Name(`ltuid`))
		if errRead != nil {
			fmt.Println("Error reading cookies from", cookieStore.Browser(), errRead)
			continue
		}
		if len(resultLtuid) <= 0 {
			fmt.Println("No cookies found in", cookieStore.Browser())
			continue
		}
		fmt.Println("Found", len(resultToken), "cookies")
		token = append(token, resultToken...)
		ltuid = append(ltuid, resultLtuid...)
		isFoundCredential = true
		break
	}
	if !isFoundCredential {
		return fmt.Errorf("account credential not found, please login to hoyolab once in Chrome/Firefox/Opera/Safari")
	}
	HoyolabToken = token
	HoyolabLtuid = ltuid
	return nil
}
