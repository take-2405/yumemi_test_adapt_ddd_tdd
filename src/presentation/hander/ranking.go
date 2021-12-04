package hander

import (
	"errors"
	"log"
	"os"
	"yumemi-coding-test-practice/src/application"
	"yumemi-coding-test-practice/src/presentation/response"
)

type RankingHandler interface {
	HandlePrintRanking()
}

type rankingHandler struct {
	rankingUseCase application.RankingUseCase
}

func NewRankingHandler(ru application.RankingUseCase) RankingHandler {
	return &rankingHandler{
		rankingUseCase: ru,
	}
}

func (rh rankingHandler) HandlePrintRanking() {

	filePath, err := parseArgment()
	if err != nil {
		log.Fatal(err)
	}

	playData, ascendingScore, err := rh.rankingUseCase.PrintScoreRanking(filePath)
	if err != nil {
		log.Fatal(err)
	}
	response.PrintResult(response.TransferPlayDataToRanking(*ascendingScore, *playData))
}

func parseArgment() (string, error) {

	filePath := os.Args[1]
	if filePath == "" {
		return filePath, errors.New("no argment.")
	}
	return filePath, nil
}
