// Here everything works on cookies for the time being

package challenge

import "net/http"

const(
	DOMAIN = "localhost"
	COOKIENAME = "user"
)
/*
func ValidateUser(response http.ResponseWriter, request *http.Request) (user User) {
	cookies := request.Cookies()
	n := len(cookies) 
	i := 0
	for i < n  {
		cookie := cookies[i]
		i++
		if(cookie.Domain == DOMAIN) {
			if(cookie.Name == COOKIENAME) {
				//validate the value in the data store and return the value
				//Check if the value is present in the db here check for user object array on challenges
				user = Store.GetUserFromUID(cookie.Value)
				return 
			}
		}
	}
	return
}

func LogInWithUserId(response http.ResponseWriter, request *http.Request, userId string, password string) (user User) {
	user = dataStore.GetUserFromUserId(userId, password) 
	// FIXME Check if the statement below works or how it is supposed to be used
	if user.Id == 0 {
		SetCookies(response, request, user)
		return
	}
	return
}

func LogInWithFbId(response http.ResponseWriter, request *http.Request, fbId string) (user User) {
	user = dataStore.GetUserFromFbId(fbId) 
	// FIXME Check if the statement below works or how it is supposed to be used
	if user.Id == 0 {
		SetCookies(response, request, user)
		return 
	}
	return 
}
*/
func SetCookies(response http.ResponseWriter, request *http.Request, user User) {
	cookie := new(http.Cookie)
	cookie.Name = COOKIENAME
	cookie.Value = user.FbId
	cookie.Domain = DOMAIN //Could be taken from the flags in future as testing becomes easier
	request.AddCookie(cookie)
}

func LogOut(response http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = COOKIENAME
	cookie.Value = "-1"
	cookie.Domain = DOMAIN //Could be taken from the flags in future as testing becomes easier
	request.AddCookie(cookie)
}
