package util

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var ParamNames *ParamName

type ParamName struct {
	ID          string
	TIME        string
	FBID        string
	PUID        string
	USERNAME    string
	NAME        string
	PASSWORD    string
	EMAIL       string
	CHALLENGEID string
	STAGEID     string
	VALIDDAYS   string
	MINCASH     string
	CASH        string
	ATTEMPTS    string
	UCID        string
	MESSAGE     string
        FBFRIENDS   string
}

func NewParamName() *ParamName {
	if ParamNames != nil {
		return ParamNames
	} else {
		ParamNames = &ParamName{
			ID:          "id",
			TIME:        "time",
			FBID:        "fbid",
			PUID:        "puid",
			USERNAME:    "uname",
			NAME:        "name",
			PASSWORD:    "pwd",
			EMAIL:       "email",
			CHALLENGEID: "cid",
			VALIDDAYS:   "vdays",
			STAGEID:     "sid",
			MINCASH:     "mcsh",
			CASH:        "csh",
			ATTEMPTS:    "atmts",
			UCID:        "ucid",
			MESSAGE:     "msg",
                        FBFRIENDS:   "fbfrnds",
		}
	}
	return ParamNames
}

const (
	BET_NEW     = "NEW"
	BET_UPDATED = "UPDATED"
	BET_LOST    = "LOST"
	BET_WON     = "WON"
)

// Function to do encrption, Can be used in generation of UID's
func Encrypt(id string) (encrypted string) {
	key := []byte(id)
	c := hmac.New(md5.New, key)
	encrypted = fmt.Sprintf("%x", c.Sum(nil))
	log.Println("Generating Key", encrypted, "for id", id)
	return
}

func GetInt(id string) (i int) {
	i, err := strconv.Atoi(id) //Both are of diff types
	if err != nil {
		log.Println("Error : ", id, "to int", err)
	}
	return
}

func GetFloat(id string) (f float64) {
	f, err := strconv.ParseFloat (id, 3) //Both are of diff types
	if err != nil {
		log.Println("Error : ", id, "to float", err)
	}
	return
}

func ParseForm(url string, key string) (value string) {
	value = ""
	urls := strings.Split(url, "?")
	parameters := strings.Split(urls[1], "&")
	for i := 0; i < len(parameters); i++ {
		arguments := strings.Split(parameters[i], "=")
		if arguments[0] == key {
			value = arguments[1]
		}
	}
	return
}
