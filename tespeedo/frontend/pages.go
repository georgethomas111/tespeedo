package frontend

import (
	"appengine"
	"challenge"
	"fmt"
	"net/http"
	"util"
)

const (
	USERALREADYEXISTS       string = "User Already Exists, Try logging In"
	USERLOGINSUCCESSFUL     string = "Login Successful"
	USERLOGINFAILED         string = "Wrong Username Password Combination"
	LOGOUTSUCCESSFUL        string = "You Have Successfully Logged Out. Hope To See You Soon"
	NOTENOUGHCASH           string = "Not Enough Cash To Participate"
	USERNOTELIGIBLEFORSTAGE string = "NULL"
	ALREADYPARTICIPATED     string = "Already Participated in the challenge"
)

// List Of templates to be in memory 
// const (
// )

func aboutHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.About.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}

func contactHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.Contact.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}

func helpHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.Help.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}

func howToPlayHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.HowToPlay.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}

func channelHandler(response http.ResponseWriter, request *http.Request) {

	webTemplates.Channel.Execute(response, nil)
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	c := appengine.NewContext(request)
	user := ValidateUser(response, request)
	c.Infof("User :", user)
	stageId := request.FormValue(paramName.STAGEID)
	c.Infof("Stage Id is :" + stageId)
	if user.Stage != 0 {
		userPresentHandler(response, request, true, user, "")
	}
	webTemplates.Header.Execute(response, *user)

	i := NewIndexView()
	i.User = user
	err := webTemplates.Index.Execute(response, i)
	if err != nil {
		c.Infof("Error Loading the index page", err)
	}
	webTemplates.Footer.Execute(response, *user)
}

func logInWithFbHandler(response http.ResponseWriter, request *http.Request) {

	c := appengine.NewContext(request)
	fbId := request.FormValue(paramName.FBID)
	c.Infof("Checking for fbid : ", fbId)

	present, user := store.LogInWithFbId(c, fbId)
	if !present {
		c.Infof("signingUp The fb Id", fbId)
		signUpWithFbHandler(response, request)
	} else {
		userPresentHandler(response, request, present, user, "")
	}
}

func logOutHandler(response http.ResponseWriter, request *http.Request) {
	LogOut(response, request)
	http.Redirect(response, request, "/", http.StatusFound)
}

func signUpWithFbHandler(response http.ResponseWriter, request *http.Request) {
	getFbId := request.FormValue(paramName.FBID)
	email := request.FormValue(paramName.EMAIL)
	name := request.FormValue(paramName.NAME)
	user := challenge.NewUser()
	user.FbId = getFbId
	user.Name = name
	user.Email = email
	user.UID = util.Encrypt(email)
	user.Key = email
	c := appengine.NewContext(request)
	store.AddUser(c, user)
	c.Infof("Signed Up The User :")
	logInWithFbHandler(response, request)
}

// Keeping Game and Stats together
const (
	FREEELIGIBLE  = 6
	RANKLISTLIMIT = 10
	NA            = 9998
	NE            = 9999
)

func gameHandler(response http.ResponseWriter, request *http.Request) {
	user := ValidateUser(response, request)
	if user.Stage == 0 {
		// Set cookies to the url to redirect.
		SetRedirectCookie(response, request)
		http.Redirect(response, request, "/", http.StatusFound)
	}
	c := appengine.NewContext(request)
	stageId := request.FormValue(paramName.STAGEID)
	message := request.FormValue(paramName.MESSAGE)
	if stageId != "" {
		err := webTemplates.Header.Execute(response, user)
		if err != nil {
			c.Infof("pages.go: gameHandler : Error while executing Header Template", err)
		}
		if user.Stage > 0 {
			g := NewGameView()
			g.User = user
			stageId := util.GetInt(request.FormValue(paramName.STAGEID))
			g.NextStage = stageId
			g.PrevStage = stageId
			g.CurrentStage = stageId
			g.Message = message
			maxEligible := 0
			if user.Stage < FREEELIGIBLE {
				maxEligible = FREEELIGIBLE
			} else {
				maxEligible = user.Stage
			}
			if stageId <= maxEligible && stageId > 0 {
				g.Stage = store.GetStage(c, stageId)
			}
			if stageId < maxEligible {
				g.NextStage = stageId + 1
			}
			if stageId > 1 {
				g.PrevStage = stageId - 1
			}

			c.Infof("Log : GameView is  -- >", g)
			err := webTemplates.Game.Execute(response, g)
			if err != nil {
				c.Infof("Could't Execute with param datastructure", err)
			}
		}
		webTemplates.Footer.Execute(response, user)
	} else {
		if user.Stage != 0 {
			userPresentHandler(response, request, true, user, message)
		}
	}
}

//Store UserId, 
func startHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	if user != nil {

		ClockStore.Add(user)
	}
}

func getStagePrev(stage *challenge.Stage) string {

	conStage := stage.Value
	minLength := 10
	if len(conStage) < minLength {
		minLength = len(conStage) - 1
	}
	return conStage[0:len(conStage)]
}

func privacyHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.Privacy.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}

// TODO Split into 3 function's
func statsSelect(request *http.Request, statsView *StatsView, user *challenge.User) {

	statsView.Selection = true
	c := appengine.NewContext(request)
	for i := 1; i <= user.Stage; i++ {

		Value := fmt.Sprintf("%d", i) + ", " + getStagePrev(store.GetStage(c, i))
		stageVal := StageInfo{
			Id:    fmt.Sprintf("%d", i),
			Value: Value,
		}
		statsView.StageData = append(statsView.StageData, stageVal)
	}
}

func statsProcess(request *http.Request, statsView *StatsView, friendData string, user *challenge.User, stageId int, time float64) {

	c := appengine.NewContext(request)
	c.Infof("Processing Statistics")
	stage := store.GetStage(c, stageId)
	if time > 100 {
		statsView.ShowSpeed = false
		time = 0
	}
	if time != 0 {
		updateScore(c, user, stage, time)
		statsView.Time = fmt.Sprintf("%f", time)
		if stage.AvgScore-time > 0 {
			statsView.Status = "Passed"
		} else {
			statsView.Status = "Failed"
		}
		speed := time / float64(len(stage.Value))
		c.Infof("LOG: Length of stage -->", len(stage.Value), "Time -- > ", time)
		statsView.Speed = fmt.Sprintf("%f", speed)
		c.Infof("Speed ->", speed)
	}

	myRank, rankData := HandleFbFriends(c, friendData, stageId, user)
	c.Infof("RankData is --> ", rankData, time)
	if rankData != nil {
		statsView.RankData = NewStatsRankData(rankData)
		c.Infof("Got RankData", statsView.RankData)
	}
	statsView.Name = user.Name
	statsView.TargetScore = fmt.Sprintf("%f", stage.AvgScore)
	statsView.MyRank = myRank
	c.Infof("My Rank is ", statsView.MyRank)

	statsView.StagePrev = getStagePrev(stage)
}

const (
	MAXDELAY = 4
)

func selectTime(cStoreTime float64, time float64) (time1 float64) {

	if cStoreTime-time < MAXDELAY {
		time1 = time
	} else {
		time1 = cStoreTime
	}
	return
}

func statsHandler(response http.ResponseWriter, request *http.Request) {

	var err error
	c := appengine.NewContext(request)
	statsView := NewStatsView()
	user := ValidateUser(response, request)
	c.Infof("Generating stats for", user)
	stageId := util.GetInt(request.FormValue(paramName.STAGEID))
	if stageId == 0 {
		statsSelect(request, statsView, user)
	} else {

		c.Infof("Stage Is is not zero " + string(stageId))
		clientTime := util.GetFloat(request.FormValue(paramName.TIME))
		cStoreTime := ClockStore.Get(user)
		time := selectTime(cStoreTime, clientTime)
		c.Infof("client Time =", clientTime, "Server time", cStoreTime)
		statsView.StageId = request.FormValue(paramName.STAGEID)
		friendData := request.FormValue(paramName.FBFRIENDS)
		c.Infof("FriendData", friendData)
		statsView.Selection = false
		statsProcess(request, statsView, friendData, user, stageId, time)
	}

	webTemplates.Header.Execute(response, user)
	err = webTemplates.Stats.Execute(response, statsView)
	if err != nil {
		c.Infof("Error Executing", err)
	}
	webTemplates.Footer.Execute(response, user)
}

func userPresentHandler(response http.ResponseWriter, request *http.Request, present bool, user *challenge.User, message string) {

	c := appengine.NewContext(request)
	if present {
		user.UID = util.Encrypt(user.Email)
		c.Infof("The user is like this now -->", user)
		store.SaveUser(c, user)
		c.Infof("User is Present", user.UserName)
		c.Infof("Modifying UID", user.UID)
		SetCookies(response, request, user)
		stageId := "-1"
		status, stageId := ReturnIfCookie(response, request)
		if status {

			ClearRedirectCookie(response, request)
		} else {

			stageId = fmt.Sprintf("%d", user.Stage)
		}

		gameUrl := fmt.Sprintf("/game?&%s=%s", paramName.STAGEID, stageId)
		http.Redirect(response, request, gameUrl, http.StatusFound)
	} else {
		fmt.Fprintf(response, USERLOGINFAILED)
	}
}

func trialHandler(response http.ResponseWriter, request *http.Request) {

	user := ValidateUser(response, request)
	webTemplates.Header.Execute(response, user)
	webTemplates.Trial.Execute(response, nil)
	webTemplates.Footer.Execute(response, user)
}
