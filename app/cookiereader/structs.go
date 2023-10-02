package cookiereader

import (
	"time"

	"github.com/zellyn/kooky"
)

type Configuration struct {
	Browser    string `json:"browser"`
	CookiePath string `json:"cookie_path"`
}

type CheckInCookie struct {
	Ltuid  kooky.Cookie `json:"ltuid"`
	Token  kooky.Cookie `json:"token"`
	Expire int64        `json:"expire"`
}

func (c CheckInCookie) IsExpired() bool {
	return c.Expire < time.Now().Unix()
}
