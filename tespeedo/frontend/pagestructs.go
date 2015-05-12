package frontend

import (
	"challenge"
	"fmt"
	"util"
)

/**
 * 
 * @Author George
 * This file contains the structures that are used in the page for parsing the contents
 *
 */

/**
 *
 * @Rating - TODO
 *
 */
type IndexView struct {
	ParamName *util.ParamName
	User      *challenge.User
}

func NewIndexView() (i *IndexView) {
	p := util.NewParamName()
	i = &IndexView{
		ParamName: p,
	}
	return
}

type LView struct {
	FbStatus  int
	ParamName *util.ParamName
}

func NewLView() *LView {
	p := util.NewParamName()
	l := &LView{
		ParamName: p,
		FbStatus:  0,
	}
	return l
}

func PViewHeader() (header []string) {
	header = append(header, "Name")
	header = append(header, "CashInBet")
	header = append(header, "Attempts")
	header = append(header, "MaxNoOfAttempts")
	header = append(header, "Rem Days")
	header = append(header, "Stage")
	return
}

type ChallengePostView struct {
	ParamName *util.ParamName
	StageId   int
	Source    string
}

func NewChallengePostView() (c *ChallengePostView) {
	p := util.NewParamName()
	c = &ChallengePostView{
		ParamName: p,
		Source:    "",
	}
	return
}

type ScoreView struct {
	ParamName *util.ParamName
	BestScore challenge.BestScore
}

func NewScoreView() (s *ScoreView) {
	p := util.NewParamName()
	s = &ScoreView{
		ParamName: p,
	}
	return
}

type GameView struct {
	ParamName    *util.ParamName
	HasChallenge bool
	Message      string
	PrevStage    int
	CurrentStage int
	NextStage    int
	Stage        *challenge.Stage
	User         *challenge.User
}

func NewGameView() (s *GameView) {
	p := util.NewParamName()
	s = &GameView{
		ParamName:    p,
		HasChallenge: false,
	}
	return
}

type DashBoardView struct {
	ParamName  *util.ParamName
	User       *challenge.User
	MYACC      bool
	Stages     []*challenge.Stage
	BestScores []challenge.BestScore
}

func NewDashBoardView() (d *DashBoardView) {
	p := util.NewParamName()
	d = &DashBoardView{
		ParamName: p,
		MYACC:     true,
	}
	return
}

type Console struct {
	Name        string
	StageId     string
	TargetScore string
	Speed       string
	Points      string
	Winner      string
	RemChances  string
}

func NewConsole() (c *Console) {
	c = &Console{}
	return
}

type StageInfo struct {
	Id    string
	Value string
}

type StatsRankData struct {
	Name       string
	Uid        string
	Pic_Square string
	Score      string
	Rank       int
}

func getScore(score float64) (resp string) {

	if score == NA {
		resp = "N.A"
	} else if score == NE {
		resp = "N.E"
	} else {
		resp = fmt.Sprintf("%f", score)
	}

	return
}

func NewStatsRankData(rankData []*Rank) []StatsRankData {

	fbRank := []StatsRankData{}
	for _, rData := range rankData {
		s := StatsRankData{
			Name:       rData.Name,
			Uid:        rData.Uid,
			Pic_Square: rData.Pic_Square,
			Score:      getScore(rData.Score),
			Rank:       rData.Rank,
		}

		fbRank = append(fbRank, s)
	}

	return fbRank
}

type StatsView struct {
	Name        string
	StageId     string
	Time        string
	TargetScore string
	Speed       string
	StagePrev   string
	Status      string
	RankData    []StatsRankData
	MyRank      int
	Selection   bool
	ShowSpeed   bool
	StageData   []StageInfo
	ParamName   *util.ParamName
}

func NewStatsView() (s *StatsView) {
	s = &StatsView{
		ParamName: util.NewParamName(),
		ShowSpeed: true,
	}
	return
}
