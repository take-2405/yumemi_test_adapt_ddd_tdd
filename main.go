package main

import (
	"log"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/infrastructure"
	"yumemi-coding-test-practice/src/presentation"
)

func main() {
	csv, err := infrastructure.ReadLogData("./game_score_log.csv")
	if err != nil {
		log.Fatal(err)
	}
	header, err := csv.Read()
	if err != nil {
		log.Fatal(err)
	}
	if !infrastructure.CheckLogHeader(header) {
		if err != nil {
			log.Fatal(err)
		}
	}
	playDatas, err := infrastructure.ParseLogData(csv)
	if err != nil {
		log.Fatal(err)
	}

	var playerScores domain.PlayerScores
	playerScores.CalcAgerageScore(playDatas)
	topScores := playerScores.SortRankingScore()

	var ranking presentation.RankingDatas
	ranking.TransferPlayDataToRanking(*topScores, playerScores)
	ranking.PrintResult()
}
