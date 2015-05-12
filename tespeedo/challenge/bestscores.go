package challenge

type BestScore struct {
	Key string
	UID string
	Score float64
	Stage int
	NoOfAttempts int
	Rank int
}

func NewBestScore () (b *BestScore) {
	b = &BestScore {
		Score : 0,
		Stage : 0,
		NoOfAttempts : 0,
	}
	return
}


func (b BestScore) GetKey () string {

	return b.Key
}

func (b BestScore) SetKey (keyVal string) {

	b.Key = keyVal
}
