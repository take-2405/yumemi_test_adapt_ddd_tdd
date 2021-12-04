package application

import (
	"sort"
	"yumemi-coding-test-practice/src/domain"
)

func CalcRanking() {

}

func SortRankingScore(p *domain.PlayerScores) *[]int {
	keys := make([]int, len(p.PlayerScore), len(p.PlayerScore))
	i := 0
	for _, score := range p.PlayerScore {
		keys[i] = score
		i++
	}
	sort.Ints(keys)
	return &keys
}
