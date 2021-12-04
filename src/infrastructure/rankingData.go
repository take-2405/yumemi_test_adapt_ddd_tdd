package infrastructure

import (
	"sort"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/domain/repository"
)

type rankingDataPersistence struct {
}

func NewRankingDataPersistence() repository.RankingDataRepository {
	return &rankingDataPersistence{}
}

func (r rankingDataPersistence) CalcPlayerAverageScore(datas domain.PlayDatas) *domain.PlayerScores {
	var PlayerScore domain.PlayerScores
	PlayerScore.PlayerScore = make(map[string]int, 500)
	for playerId, playData := range datas.Data {
		score := playData.Score / playData.Count
		PlayerScore.PlayerScore[playerId] = score
	}
	return &PlayerScore
}

func (r rankingDataPersistence) GetAscendingOrderScore(playerScore *domain.PlayerScores) *[]int {
	keys := make([]int, len(playerScore.PlayerScore), len(playerScore.PlayerScore))
	i := 0
	for _, score := range playerScore.PlayerScore {
		keys[i] = score
		i++
	}
	sort.Ints(keys)
	return &keys
}
