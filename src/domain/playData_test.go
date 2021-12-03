package domain

import (
	"log"
	"os"
	"reflect"
	"testing"
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

//func TestPlayerScores_CalcAgerageScore(t *testing.T) {
//	var playerScores PlayerScores
//	type fields struct {
//		PlayerScore map[string]int
//	}
//	type args struct {
//		playDatas *infrastructure.PlayDatas
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//		{
//			name:   "SuccessCalcAgerageScore",
//			fields: fields{playerScores.PlayerScore},
//			args:   args{playData},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &PlayerScores{
//				PlayerScore: tt.fields.PlayerScore,
//			}
//			r.CalcAgerageScore(tt.args.playDatas)
//		})
//	}
//}

func TestPlayerScores_SortRankingScore(t *testing.T) {

	var playerScores PlayerScores
	playerScores.CalcAgerageScore(playData)
	type fields struct {
		PlayerScore map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]int
	}{
		// TODO: Add test cases.
		{
			name:   "SuccessRankingDatas_SortRankingScore",
			fields: fields{playerScores.PlayerScore},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PlayerScores{
				PlayerScore: tt.fields.PlayerScore,
			}
			if got := r.SortRankingScore(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortRankingScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
