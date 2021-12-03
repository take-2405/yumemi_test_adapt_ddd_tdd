package domain

import (
	"encoding/csv"
	"reflect"
	"testing"
	"yumemi-coding-test-practice/src/infrastructure"
)

func TestParseLogData(t *testing.T) {
	record, _ := infrastructure.ReadLogData("./../../game_score_log.csv")
	record.Read()
	type args struct {
		csv *csv.Reader
	}
	tests := []struct {
		name string
		args args
		want *PlayDatas
	}{
		// TODO: Add test cases.
		{
			name: "SuccessParseCSVData",
			args: args{record},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLogData(tt.args.csv); reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLogData() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestRankingDatas_CalcAgerageScore(t *testing.T) {
//	var rankingData RankingDatas
//	record, _ := infrastructure.ReadLogData("./../../game_score_log.csv")
//	record.Read()
//	playData := ParseLogData(record)
//
//	type fields struct {
//		Datas map[int]string
//	}
//	type args struct {
//		data *PlayDatas
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//		{
//			name:   "SuccessCalcAgerageScore",
//			fields: fields{rankingData.Datas},
//			args:   args{playData},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &RankingDatas{
//				Datas: tt.fields.Datas,
//			}
//			r.CalcAgerageScore(tt.args.data)
//		})
//	}
//}

func TestRankingDatas_SortRankingScore(t *testing.T) {
	var rankingData RankingDatas
	record, _ := infrastructure.ReadLogData("./../../game_score_log.csv")
	record.Read()
	playData := ParseLogData(record)
	rankingData.CalcAgerageScore(playData)

	type fields struct {
		Datas map[int]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]int
	}{
		// TODO: Add test cases.
		{
			name:   "SuccessRankingDatas_SortRankingScore",
			fields: fields{rankingData.Datas},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RankingDatas{
				Datas: tt.fields.Datas,
			}
			if got := r.SortRankingScore(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortRankingScore() = %v, want %v", got, tt.want)
			}
		})
	}
}