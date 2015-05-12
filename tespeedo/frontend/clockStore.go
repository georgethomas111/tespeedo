package frontend

import ( 
    "challenge"
    "time"
    "log"
)

var ClockStore* Clocks = NewClocks ()

type Clocks struct {
     
     Store map[int]time.Time // Mapping uid and base time
}

func NewClocks () (c *Clocks) {
    c = &Clocks {
        Store : make (map [int]time.Time), 
    } 
    return
}

func (c* Clocks) Add (user *challenge.User) {
    c.Store [user.Id] = time.Now ()
    log.Println ("ClockStore : Adding Time ", c.Store [user.Id])
}

func (c* Clocks) Get (user *challenge.User) float64 {
    endTime   := time.Now ()
    startTime := c.Store [user.Id]
    log.Println ("ClockStore : End Time ", endTime)
    delete (c.Store, user.Id)
    return endTime.Sub (startTime).Seconds () 
} 
