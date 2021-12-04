package infrastructure

import (
	"encoding/csv"
	"reflect"
	"testing"
)

func TestParseLogData(t *testing.T) {
	record, _ := ReadLogData("./../../game_score_log.csv")
	record.Read()
	type args struct {
		csv *csv.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *PlayDatas
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "SuccessParseCSVData",
			args:    args{record},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ParseLogData(tt.args.csv)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLogData() = %v, want %v", got, tt.want)
			}
			if gotErr != tt.wantErr {
				t.Errorf("ReadLogData is  error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
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
