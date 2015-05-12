package frontend

import (
	"challenge"
        "log"
        "appengine"
        "encoding/json" 
)

func updateScore(c appengine.Context, user *challenge.User, stage *challenge.Stage, time float64) {
	if time < stage.AvgScore {
		if user.Stage <= stage.Id {
			user.Stage = stage.Id + 1
			c.Infof ("Stage Id", stage.Id)
			c.Infof ("updating users stage to", user.Stage)
			defer store.SaveUser(c, user)
		}
	} else {
		log.Println("Did'nt pass stage", time)
	}

	bestScore := store.GetBestScore (c, user.UID, stage.Id)

	if bestScore.Stage == 0 {
		bestScore.UID = user.UID
		bestScore.Stage = stage.Id
		bestScore.Score = time
		bestScore.NoOfAttempts++
		store.AddBestScore(c, bestScore)
	} else if bestScore.Score > time {
		bestScore.NoOfAttempts++
		bestScore.Score = time
		store.SaveBestScore(c, bestScore)
	} else {
		bestScore.NoOfAttempts++
		store.SaveBestScore(c, bestScore)
	}

	if time < stage.AvgScore*4 {
		stage.UpdateTime(time)
		defer store.SaveStage(c, stage)
	}
}

type Rank struct {
	Name       string
	Uid        string
	Pic_Square string
	Score      float64
	Rank       int

	LeftRank  *Rank
	RightRank *Rank
}

func NewRank(fbData FbData) (rank *Rank) {
	rank = &Rank{
		Name:      fbData.Name,
		Uid:       fbData.Uid,
		Pic_Square:fbData.Pic_Square,
		Score:     fbData.Score,
		Rank:      -1,
		LeftRank:  nil,
		RightRank: nil,
	}
	return
}


type RankList struct {
	MyRank *Rank
}

var RankChannel = make (chan *Rank, 100)

//Keeping max in queue to be 100
func NewRankList(myRank *Rank) (rankList *RankList) {
	rankList = &RankList{
		MyRank: myRank,
	}
	return
}

func Insert(insNode *Rank, rank *Rank) (*Rank) {

	if insNode == nil {
		insNode = rank
	} else if rank.Score > insNode.Score {
		insNode = Insert(insNode.RightRank, rank)
	} else {
		insNode = Insert(insNode.LeftRank, rank)
	}
	return insNode
}

// Returns the parent after adding the children to it.
func AddRank (c appengine.Context, parent* Rank, child* Rank) (*Rank) {
	if parent == nil {
            return child
	}
	if child.Score <= parent.Score {
		parent.LeftRank = AddRank (c, parent.LeftRank, child)
	} else {
		parent.RightRank = AddRank (c, parent.RightRank, child)
	}
	return parent
}

// Add below myRank always
func AddToRankList(c appengine.Context, rankList *RankList, rank *Rank) (*RankList) {
	rankList.MyRank = AddRank (c, rankList.MyRank, rank)
	return rankList
}

func GetSorted (c appengine.Context, rank* Rank) {
	if rank == nil {
		return
	} else {
		GetSorted (c, rank.LeftRank)
		c.Infof ("Sorted:" , rank.Name, ":", rank.Score)
		RankChannel <- rank
		GetSorted (c, rank.RightRank)
	}
}

/**
*
* Limit the number sent to 10
**/

func GetRankSlice (c appengine.Context, rankList *RankList, limit int) (list []*Rank) {
	for i := 0 ; i < limit; i++ {
		rank := <-RankChannel
		rank.Rank = i + 1 // Assigning the right rank
		c.Infof ("Rank : Name" , rank.Rank, ":", rank.Name)
		list = append (list, rank)
	}
	return list
}

type FbData struct {
   Name       string
   Uid        string
   Pic_Square string
   Score      float64
   Rank       int
}


func sanitizeScore (score float64,
user *challenge.User,
stageId int) (float64) {
	maxEligible := user.Stage
	if user.Stage < FREEELIGIBLE {
		maxEligible = FREEELIGIBLE
	}
	if maxEligible < stageId {
		 score = NE
	} else if score == 0 {
		 score = NA
	}
	return score
}

/*
getRankList returns a max RANKLISTLIMIT number of element of the sorted ranklist
*/
func getRankList (c appengine.Context,
fbData *[]FbData,
stageId int,
user *challenge.User) (int, []*Rank) {

	noOfFriends := len (*fbData)
	c.Infof ("NoOfFriends = ",  noOfFriends)
	myScore := store.GetBestScore (c, user.UID, stageId)
	(*fbData)[noOfFriends-1].Score = sanitizeScore (myScore.Score, user, stageId)
	myRank := NewRank ((*fbData)[noOfFriends-1])
	rankList := NewRankList (myRank)
	notPresent := 0
	for i := 0; i < noOfFriends - 1 ; i++ {
		present, user := store.LogInWithFbId (c, (*fbData)[i].Uid)
		if present {
			bestScore  := store.GetBestScore (c, user.UID, stageId)
			log.Println ("Game.go: ", bestScore)
			c.Infof ("Game.go: StageId = ", stageId, "UserName =", user.Name,"bestScore", bestScore.Score)
			(*fbData)[i].Score = sanitizeScore (bestScore.Score, user, stageId)
			rank     := NewRank((*fbData)[i])
			rankList  = AddToRankList(c, rankList, rank)
		} else {
		   notPresent ++;
		}
	}
	go GetSorted (c, rankList.MyRank)
	sortedData := GetRankSlice (c, rankList, noOfFriends - notPresent)
	c.Infof ("Size of sorted Data is", len (sortedData))
	if (len (sortedData) >= RANKLISTLIMIT) {
		sortedData = sortedData [0:RANKLISTLIMIT-1]
	}
	return rankList.MyRank.Rank, sortedData
}

/**
* Use the dataStore to get the value.
* if value is greater then use other mechanisms.
*/
func HandleFbFriends (c appengine.Context,
                      friendData string, stageId int,
                      user *challenge.User) (int, []*Rank) {
      var decoded []FbData
      c.Infof ("The Data is :", friendData)
      err := json.Unmarshal ([]byte(friendData), &decoded)
      if err == nil {
           myRank, rankList := getRankList (c, &decoded, stageId, user)
           return myRank, rankList
      } else {
           c.Infof ("Error Unmarshalling", err)
           return -1, nil
      }
      return -1, nil
}
