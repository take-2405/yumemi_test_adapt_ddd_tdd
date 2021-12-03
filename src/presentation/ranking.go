package presentation

import (
	"yumemi-coding-test-practice/src/domain"
)

type RankingData struct {
	Rank   int
	Player string
	Score  int
}

type RankingDatas struct {
	datas []RankingData
}

func (r *RankingDatas) TransferPlayDataToRanking(topScores []int, playerScores domain.PlayerScores) {
	var rankingData RankingData
	rank := 1
	beforeScore := 0
	for i := len(topScores) - 1; i > 0; i-- {
		if topScores[i] != beforeScore {
			for playerId, playerScore := range playerScores.PlayerScore {
				if topScores[i] == playerScore {
					rankingData.Rank = rank
					rankingData.Player = playerId
					rankingData.Score = topScores[i]
					r.datas = append(r.datas, rankingData)
				}
			}
			rank++
		}
		if rank == 11 {
			break
		}
		beforeScore = topScores[i]
	}
}
