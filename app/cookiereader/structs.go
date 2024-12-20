package cookiereader

import (
	"slices"
	"time"

	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
)

// CookieData is a struct to store cookies key format <browser>:<profile>
type CookieData struct {
	Name             string    `json:"name"`
	Value            string    `json:"value"`
	Quoted           bool      `json:"quoted"`
	Path             string    `json:"path"`
	Domain           string    `json:"domain"`
	Expires          time.Time `json:"expires"`
	MaxAge           int       `json:"max_age"`
	Secure           bool      `json:"secure"`
	HttpOnly         bool      `json:"http_only"`
	SameSite         int       `json:"same_site"`
	Partitioned      bool      `json:"partitioned"`
	Creation         time.Time `json:"creation"`
	Browser          string    `json:"browser"`
	Profile          string    `json:"profile"`
	IsDefaultProfile bool      `json:"is_default_profile"`
	FilePath         string    `json:"file_path"`
}

type CheckInCookieV2 struct {
	// Profile string         `json:"profile"`
	// Browser string         `json:"browser"`
	Cookies []CookieData `json:"cookies"`
	Expire  int64        `json:"expire"`
}

func (c CheckInCookieV2) IsExpired() bool {
	return c.Expire < time.Now().Unix()
}

func (c CheckInCookieV2) IsEligbleForCheckinZZZ() bool {
	requireCookieCount := 0
	for _, cookie := range c.Cookies {
		if slices.Contains(myconsants.CookieNamesForZZZ, cookie.Name) {
			requireCookieCount++
		}
	}
	return requireCookieCount == len(myconsants.CookieNamesForZZZ)
}
