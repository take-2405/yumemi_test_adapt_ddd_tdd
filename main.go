package main

import (
	"log"
	"yumemi-coding-test-practice/src/application"
	"yumemi-coding-test-practice/src/infrastructure"
	"yumemi-coding-test-practice/src/presentation"
)

func main() {
	arg := presentation.ParseArgment()
	csv, err := infrastructure.ReadLogData(arg.FilePath)
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

	playerScores := playDatas.CalcAgerageScore()

	topScores := application.SortRankingScore(playerScores)

	var ranking presentation.RankingDatas
	ranking.TransferPlayDataToRanking(*topScores, *playerScores)
	ranking.PrintResult()
}
