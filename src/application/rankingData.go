package application

import (
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/domain/repository"
)

type RankingUseCase interface {
	PrintScoreRanking(filePath string) (*domain.PlayerScores, *[]int, error)
}

type rankingUseCase struct {
	rankingRepository  repository.RankingDataRepository
	playDataRepository repository.PlayDataRepository
}

func NewRankingUseCase(p repository.PlayDataRepository, r repository.RankingDataRepository) RankingUseCase {
	return &rankingUseCase{
		rankingRepository:  r,
		playDataRepository: p,
	}
}

// usecaseを生成
func (ru rankingUseCase) PrintScoreRanking(filePath string) (*domain.PlayerScores, *[]int, error) {
	csv, err := ru.playDataRepository.ReadPlayData(filePath)
	if err != nil {
		return nil, nil, err
	}

	playData, err := ru.playDataRepository.ParsePlayData(csv)
	if err != nil {
		return nil, nil, err
	}

	playerScore := ru.rankingRepository.CalcPlayerAverageScore(playData)
	ascendingScore := ru.rankingRepository.GetAscendingOrderScore(playerScore)

	return playerScore, ascendingScore, nil
}
