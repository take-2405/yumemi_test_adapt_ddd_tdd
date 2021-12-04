package infrastructure

import (
	"log"
	"reflect"
	"testing"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/domain/repository"
)

var playDatas domain.PlayDatas

func Test_rankingDataPersistence_CalcPlayerAverageScore(t *testing.T) {
	var playDatas domain.PlayDatas

	playData := NewPlayDataPersistence()
	playDatarepo := repository.PlayDataRepository(playData)
	csv, err := playDatarepo.ReadPlayData("./../../game_score_log.csv")
	if err != nil {
		log.Fatal(err)
	}
	_, err = csv.Read()
	if err != nil {
		log.Fatal(err)
	}
	playDatas, err = playDatarepo.ParsePlayData(csv)
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		datas domain.PlayDatas
	}
	tests := []struct {
		name string
		args args
		want *domain.PlayerScores
	}{
		{
			name: "successCalcPlayerAverageScore",
			args: args{datas: playDatas},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rankingDataPersistence{}
			if got := r.CalcPlayerAverageScore(tt.args.datas); reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcPlayerAverageScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rankingDataPersistence_GetAscendingOrderScore(t *testing.T) {
	var playDatas domain.PlayDatas

	playData := NewPlayDataPersistence()
	playDatarepo := repository.PlayDataRepository(playData)
	csv, err := playDatarepo.ReadPlayData("./../../game_score_log.csv")
	if err != nil {
		log.Fatal(err)
	}
	_, err = csv.Read()
	if err != nil {
		log.Fatal(err)
	}
	playDatas, err = playDatarepo.ParsePlayData(csv)
	if err != nil {
		log.Fatal(err)
	}
	r := rankingDataPersistence{}
	playerScore := r.CalcPlayerAverageScore(playDatas)
	type args struct {
		playerScore *domain.PlayerScores
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		{
			name: "successGetAscendingOrderScore",
			args: args{playerScore: playerScore},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rankingDataPersistence{}
			if got := r.GetAscendingOrderScore(tt.args.playerScore); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAscendingOrderScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
