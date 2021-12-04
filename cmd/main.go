package main

import (
	"yumemi-coding-test-practice/src/application"
	"yumemi-coding-test-practice/src/domain/repository"
	"yumemi-coding-test-practice/src/infrastructure"
	"yumemi-coding-test-practice/src/presentation/hander"
)

func main() {
	playDataPersistence := infrastructure.NewPlayDataPersistence()
	rankingDataPersistence := infrastructure.NewRankingDataPersistence()

	rankingRepository := repository.RankingDataRepository(rankingDataPersistence)
	playDataRepository := repository.PlayDataRepository(playDataPersistence)

	rankingUseCase := application.NewRankingUseCase(playDataRepository, rankingRepository)

	handler := hander.NewRankingHandler(rankingUseCase)
	handler.HandlePrintRanking()
}
