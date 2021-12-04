package repository

import "yumemi-coding-test-practice/src/domain"

type RankingDataRepository interface {
	CalcPlayerAverageScore(datas domain.PlayDatas) *domain.PlayerScores
	GetAscendingOrderScore(playerScore *domain.PlayerScores) *[]int
}
