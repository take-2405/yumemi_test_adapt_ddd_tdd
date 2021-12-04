package application

import (
	"log"
	"os"
	"reflect"
	"testing"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/infrastructure"
)

var playData *infrastructure.PlayDatas

func TestMain(m *testing.M) {
	var err error
	record, _ := infrastructure.ReadLogData("./../../game_score_log.csv")
	record.Read()
	playData, err = infrastructure.ParseLogData(record)
	if err != nil {
		log.Fatal()
	}
	os.Exit(0)
}

func TestPlayerScores_SortRankingScore(t *testing.T) {
	var playerScores domain.PlayerScores
	playData.CalcAgerageScore()

	//playData.CalcAgerageScore(&playerScores)
	type args struct {
		p *domain.PlayerScores
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		// TODO: Add test cases.
		{
			name: "SuccessRankingDatas_SortRankingScore",
			args: args{&playerScores},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortRankingScore(tt.args.p); reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortRankingScore() = %v, want %v", got, tt.want)
			}
		})
	}
}