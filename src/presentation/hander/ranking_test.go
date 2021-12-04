package hander

import (
	"flag"
	"os"
	"testing"
	"yumemi-coding-test-practice/src/application"
	"yumemi-coding-test-practice/src/domain/repository"
	"yumemi-coding-test-practice/src/infrastructure"
)

func TestMain(m *testing.M) {
	//flag.CommandLine.Set("","")
	flag.CommandLine.Set("", "./game_score_log.csv")

	os.Exit(0)
}

func Test_rankingHandler_HandlePrintRanking(t *testing.T) {
	playDataPersistence := infrastructure.NewPlayDataPersistence()
	rankingDataPersistence := infrastructure.NewRankingDataPersistence()

	rankingRepository := repository.RankingDataRepository(rankingDataPersistence)
	playDataRepository := repository.PlayDataRepository(playDataPersistence)

	rankingUseCase := application.NewRankingUseCase(playDataRepository, rankingRepository)
	type fields struct {
		rankingUseCase application.RankingUseCase
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "successHandlePrintRanking",
			fields: fields{rankingUseCase: rankingUseCase},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := rankingHandler{
				rankingUseCase: tt.fields.rankingUseCase,
			}
			rh.HandlePrintRanking()
		})
	}
}
