package frontend

// Change the name of the package to a server so that it will have the functions that are the basic requirements of any server

import (
	"challenge"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"util"
)

var (
	listenAddress = flag.String("http", ":80", "http listen address")
	hostname      = flag.String("host", "localhost", "http host name")
)

//TODO's Should add logs for future requirements
type Page struct {
	Title string
	Body  []byte
}

func loadPage(response http.ResponseWriter, filename string) (*Page, error) {

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}

func GetTemplate(loc string) *template.Template {

	myTemplate, err := template.ParseFiles(loc)
	if err != nil {
		log.Println("Parse Error :", loc, ":", err)
	}
	return myTemplate
}

func rootHandler(response http.ResponseWriter, request *http.Request) {

	fileName := request.URL.Path[1:]
	switch fileName {

	case "index":
		indexHandler(response, request)
		break
	case "":
		indexHandler(response, request)
		break
	default:
		http.NotFound(response, request)
	}
}

type MyTemplate struct {
	Header    *template.Template
	Footer    *template.Template
	About     *template.Template
	Contact   *template.Template
	Channel   *template.Template
	Game      *template.Template
	Help      *template.Template
	HowToPlay *template.Template
	Index     *template.Template
	Stats     *template.Template
	Trial     *template.Template
	Privacy   *template.Template
}

// Caching in memory for faster access
func NewMyTemplate() (myTemp *MyTemplate) {
	myTemp = &MyTemplate{
		Header:    GetTemplate("templates/header.html"),
		Footer:    GetTemplate("templates/footer.html"),
		About:     GetTemplate("templates/about.html"),
		Contact:   GetTemplate("templates/contact.html"),
		Channel:   GetTemplate("templates/channel.html"),
		Game:      GetTemplate("templates/game.html"),
		Help:      GetTemplate("templates/help.html"),
		HowToPlay: GetTemplate("templates/howtoplay.html"),
		Index:     GetTemplate("templates/index.html"),
		Stats:     GetTemplate("templates/stats.html"),
		Trial:     GetTemplate("templates/trial.html"),
		Privacy:   GetTemplate("templates/privacy.html"),
	}
	return
}

var store *challenge.TespeedoStore
var paramName *util.ParamName
var webTemplates *MyTemplate

// Do FB SignUp Later
func init() {
	store = challenge.Store
	webTemplates = NewMyTemplate()

	paramName = util.NewParamName()
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", rootHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/help", helpHandler)
	http.HandleFunc("/channel", channelHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/howToPlay", howToPlayHandler)
	http.HandleFunc("/logInWithFb", logInWithFbHandler)
	http.HandleFunc("/logOut", logOutHandler)
	http.HandleFunc("/signUpWithFb", signUpWithFbHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/trial", trialHandler)
	http.HandleFunc("/privacypolicy", privacyHandler)

	http.ListenAndServe(*listenAddress, nil)
}
