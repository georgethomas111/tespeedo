package challenge

// Define a struct to hold row data  
type Stage struct {
	Key            string
	Id             int
	Value          string
	TotalTime      float64
	NoOfEntries    int
	AvgScore       float64
}

//Form the type and call this function
//Sets average score for calculations
func (s *Stage) UpdateTime(time float64) {
	s.TotalTime += time
	s.NoOfEntries += 1
}

func (s *Stage) GetId () int {

	return s.Id
}

func (s Stage) GetKey () string {

	return string (s.Id)
}

func (s Stage) SetKey (keyVal string) {

	s.Key = keyVal
}
