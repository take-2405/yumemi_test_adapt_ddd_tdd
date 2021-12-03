package domain

import (
	"sort"
	"yumemi-coding-test-practice/src/infrastructure"
)

type PlayerScores struct {
	PlayerScore map[string]int
}

func (p *PlayerScores) CalcAgerageScore(playDatas *infrastructure.PlayDatas) {
	p.PlayerScore = make(map[string]int, 500)
	for playerId, playData := range playDatas.Datas {
		score := playData.Score / playData.Count
		p.PlayerScore[playerId] = score
	}
}

func (p *PlayerScores) SortRankingScore() *[]int {
	keys := make([]int, len(p.PlayerScore), len(p.PlayerScore))
	i := 0
	for _, score := range p.PlayerScore {
		keys[i] = score
		i++
	}
	sort.Ints(keys)
	return &keys
}
