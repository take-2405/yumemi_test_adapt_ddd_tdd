package infrastructure

import (
	"reflect"
	"testing"
	"yumemi-coding-test-practice/src/domain"
)

var playDatas domain.PlayDatas

func Test_rankingDataPersistence_CalcPlayerAverageScore(t *testing.T) {
	var playDatas domain.PlayDatas
	playDatas.Data = make(map[string]domain.PlayData)
	playDatas.Data["a"] = domain.PlayData{2, 110}
	playDatas.Data["b"] = domain.PlayData{3, 130}

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
	var playDatas domain.PlayerScores
	playDatas.PlayerScore = make(map[string]int)
	playDatas.PlayerScore["a"] = 100
	playDatas.PlayerScore["b"] = 200

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
			args: args{playerScore: &playDatas},
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
