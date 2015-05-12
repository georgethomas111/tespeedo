package challenge

import (
	"log"
	"fmt"
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
)

// The idea is simple have a unique identifier to map used to index and delete it 
type TespeedoStore struct {

}

const (
USERCACHEPREFIX = "user"
//As stage Id is int adding a negative number as prefix instead
STAGECACHEPREFIX = "stage"
BESTSCOREPREFIX = "bscore"
)

var Store* TespeedoStore = NewTespeedoStore ()
func NewTespeedoStore() *TespeedoStore {
	tStore := &TespeedoStore{
	}
	return tStore
}

// ------------------------- Put in separate file in the end

type DataStoreElem interface {

	GetKey () string
	SetKey (string)
}

func PutElem (c appengine.Context, storeName string, elem DataStoreElem) (err error) {
	if elem.GetKey () == "" {
		low, _, _ := datastore.AllocateIDs (c, storeName, nil, 1)
		elem.SetKey (fmt.Sprintf ("%d", low))
	}

	key    :=  datastore.NewKey(c, storeName, elem.GetKey (), 0, nil)
	_, err = datastore.Put (c, key, elem)
	if err != nil {
		return err
	} else {
		return err
	}
	return
}

/*

There is no need to worry about errors here.

*/
func AddToCache (c appengine.Context, key string, elem interface {}) {

	item := &memcache.Item {
		Key    : key,
		Object : elem,
	}
	memcache.Gob.Set (c, item)
	c.Infof ("Added element with key" + key, elem)
}

// -------------------------------

func (t* TespeedoStore) GetUser (c appengine.Context, key string, value interface {}) (*User) {

	user  := new (User)
	update := true
	ckey  := USERCACHEPREFIX + key + value.(string)
	_, err := memcache.Gob.Get (c, ckey, &user)
	if key != "UID" {
	     update = false
	}
	if (err == memcache.ErrCacheMiss) {

		c.Infof ("Cache Miss for key", ckey)
		query := datastore.NewQuery ("UserStore").Filter (key, value)
		for tuple := query.Run(c); ; {
			_, err := tuple.Next (user)
			if err == datastore.Done {
				break
			}
		}
		if update {
			AddToCache (c, ckey, &user)
		}
	}
	return user
}

func (t *TespeedoStore) GetUserFromUID(c appengine.Context, UID string) (*User) {

	key  := "UID="
	user := t.GetUser (c, key, UID)
	return user
}

func (t *TespeedoStore) GetUserFromPUID (c appengine.Context, PUID string) *User {
	key  := "PUID="
	user := t.GetUser (c, key, PUID)
	return user
}


func (t *TespeedoStore) GetUserFromId(c appengine.Context, id int) *User {
	key  := "ID="
	user := t.GetUser (c, key, id)
	return user
}



func (t* TespeedoStore) PutUser (c appengine.Context, user *User, key string) (err error) {
	dbKey    :=  datastore.NewKey (c, "UserStore", key, 0, nil)
	_, err    = datastore.Put (c, dbKey, user)
	return
}

/*
 Not caching this one as it is required only once
*/
func (t *TespeedoStore) GetUserFromFbId(c appengine.Context, fbId string) (User) {
	user := new (User)
	user.SetKey(fbId)

	query := datastore.NewQuery ("UserStore").Filter ("FbId = ", fbId)
	for tuple := query.Run(c); ; {
		_, err := tuple.Next (user)
		if err == datastore.Done {
			break
		}
	}
	return *user
}

func (t* TespeedoStore) GetStageFromDb (c appengine.Context, key string, value interface {}) (*Stage) {

	stage  := new (Stage)
		query := datastore.NewQuery ("StageStore").Filter (key, value)
		for tuple := query.Run(c); ; {
			_, err := tuple.Next (stage)
			if err == datastore.Done {
				break
			}
		}
		log.Println ("Got Stage :", stage)
	return stage
}

func (t *TespeedoStore) GetStage(c appengine.Context, stageId int) (*Stage) {

	sKey   := STAGECACHEPREFIX + string (stageId)
	stage  := new (Stage)
	_, err := memcache.Gob.Get (c, sKey, &stage)
	if (err == memcache.ErrCacheMiss) {

		key  := "Id="
		stage = t.GetStageFromDb (c, key, stageId)
		AddToCache (c, sKey, &stage)
	}
	return stage
}

// TODO: 1: Add channels to get the work done as per scale (channels can be present here itself
//       2: Add & Save are separate the working is the same. Use one of them

func (t *TespeedoStore) AddUser(c appengine.Context, user *User) {
        key := USERCACHEPREFIX + "UID=" + user.UID
	AddToCache (c, key, user)
	PutElem (c, "UserStore", user)
}

// The function is the same saves the User
func (t *TespeedoStore) SaveUser(c appengine.Context, user *User) {
        key := USERCACHEPREFIX + "UID=" + user.UID
	AddToCache (c, key, user)
	PutElem (c, "UserStore", user)
}

func (t *TespeedoStore) SaveStage(c appengine.Context, stage *Stage) {
	PutElem (c, "StageStore", stage)
}

func (t *TespeedoStore) AddBestScore(c appengine.Context, bestScore *BestScore) {
	bKey := BESTSCOREPREFIX + string (bestScore.Stage) + bestScore.UID
	AddToCache (c, bKey, bestScore)
	PutElem (c, "BestScoreStore", bestScore)
}

func (t *TespeedoStore) SaveBestScore(c appengine.Context, bestScore *BestScore) {
	bKey := BESTSCOREPREFIX + string (bestScore.Stage) + bestScore.UID
	AddToCache (c, bKey, bestScore)
	PutElem (c, "BestScoreStore", bestScore)
}

func (t *TespeedoStore) GetBestScore(c appengine.Context, UID string, stage int) (*BestScore) {

	bKey := BESTSCOREPREFIX + string (stage) + UID
	bestScore := new (BestScore)
	_, err := memcache.Gob.Get (c, bKey, &bestScore)
	if (err == memcache.ErrCacheMiss) {
		query := datastore.NewQuery ("BestScoreStore")
		query  = query.Filter ("UID=", UID)
		query  = query.Filter ("Stage=", stage)
		log.Println ("Stage Id is", query, stage)
		query.Run(c).Next (bestScore)
		AddToCache (c, bKey, &bestScore)
        }
	return bestScore
}

func (t *TespeedoStore) GetBestScores(c appengine.Context, UID int) (scores []*BestScore) {

	return
}


func (t* TespeedoStore) GetUserFromUserName (userName string, password string) (*User) {
	var user User
	return &user
}

func (t* TespeedoStore) GetUserFromEmail (email string, password string) (user *User) {
	return
}

func (t *TespeedoStore) LogInWithUserName(username string, password string) (present bool, user *User) {
	user = t.GetUserFromUserName(username, password)

	if user.Id != 0 {
		present = true
		return
	}
	present = false
	return
}
func (t *TespeedoStore) LogInWithEmail(email string, password string) (present bool, user *User) {
	user = t.GetUserFromEmail(email, password)
	if user.Id != 0 {
		present = true
		return
	}
	present = false
	return
}

func (t *TespeedoStore) LogInWithFbId(c appengine.Context, fbId string) (present bool, user *User) {
	dbUser := t.GetUserFromFbId (c, fbId)
	user    = &dbUser
	if user.Stage > 0 {
		present = true
	} else {
		present = false
	}

	return
}
