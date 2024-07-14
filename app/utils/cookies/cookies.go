package cookies

import "github.com/gin-contrib/sessions/cookie"

var CookieStore cookie.Store

func InitCookieStore() {
	CookieStore = cookie.NewStore([]byte("secret"))
}
