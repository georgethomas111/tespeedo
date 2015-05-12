package challenge

// TODO Use coins which gets added with every appreciable work
// TODO Challenge points is not being used at present but could be used if it scales :) 
// TODO Make a standard like names of columns in db are same and make a function that can retrieve id for the type given (email , fbid)


type User struct {

	Id int
	Key string
	UID string
	PID string
	FbId string
	UserName string
	Password string
	Name string
	Email string
	Stage int
	Points int
	Cash int
}

func NewUser() (user *User) {
	user = &User {
		Stage : 1,
		Points : 100,
		Cash : 1000,
	}
	return
}

func (u* User) GetId () int {

	return u.Id        
}

func (u User) GetKey () string {

	return u.Key        
}

func (u User) SetKey (keyVal string) {

	u.Key = keyVal        
}
