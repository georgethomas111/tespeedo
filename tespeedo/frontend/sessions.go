// Here everything works on cookies for the time being

package frontend

import (
	"challenge"
	"log"
	"net/http"
        "appengine"
)

const (
	DOMAIN             = "tespeedo.com"
	UIDCOOKIENAME      = "user"
        STAGEIDCOOKIE      = "_sid"
)

// Done
func ValidateUser(response http.ResponseWriter, request *http.Request) (user *challenge.User) {
	var nullUser challenge.User
	cookies := request.Cookies()
	n := len(cookies)
	i := 0
	log.Println("cookie Length =", n)
	for i < n {
		cookie := cookies[i]
		i++
		if cookie.Name == UIDCOOKIENAME {
			//validate the value in the data store and return the value
			//Check if the value is present in the db here check for user object array on challenges
                        c   := appengine.NewContext (request)
			user = store.GetUserFromUID(c, cookie.Value)
                        log.Println ("session.go : User Obtained From Session:", user) 
		}
	}
	if user == nil {
		return &nullUser
	}
	return
}

func ReturnIfCookie (response http.ResponseWriter, request *http.Request) (bool, string) {
	
	cookies := request.Cookies()
	n := len(cookies)
	i := 0
	log.Println("cookie Length =", n)
	for i < n {
		cookie := cookies[i]
		i++
		if cookie.Name == STAGEIDCOOKIE {
			if cookie.Value != "-1" {
				return true, cookie.Value
			}
		}
	}
	return false, "-1"
}

//Done
func SetCookies(response http.ResponseWriter, request *http.Request, user *challenge.User) {
	cookie := new(http.Cookie)
	cookie.Name = UIDCOOKIENAME
	cookie.Value = user.UID
	cookie.MaxAge = 0
	cookie.Path = "/"
	log.Println("cookie added to request =", cookie)
	http.SetCookie(response, cookie)

}

func SetRedirectCookie (response http.ResponseWriter, request *http.Request) {

	cookie := new(http.Cookie)
	cookie.Name = STAGEIDCOOKIE
	cookie.Value = request.FormValue (paramName.STAGEID)
	cookie.MaxAge = 0
	cookie.Path = "/"
	log.Println("cookie added to request =", cookie)
	http.SetCookie(response, cookie)
} 

func ClearRedirectCookie (response http.ResponseWriter, request *http.Request) {

	cookie := new(http.Cookie)
	cookie.Name = STAGEIDCOOKIE
	cookie.Value = "-1"
	cookie.MaxAge = 0
	cookie.Path = "/"
	log.Println("cookie added to request =", cookie)
	http.SetCookie(response, cookie)
}

//Done
func LogOut(response http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = UIDCOOKIENAME
	cookie.Value = "-1"
	http.SetCookie(response, cookie)
}
