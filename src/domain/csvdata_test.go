package domain

import (
	"encoding/csv"
	"reflect"
	"testing"
	"yumemi-coding-test-practice/src/infrastructure"
)

func TestParseLogData(t *testing.T) {
	record,_:=infrastructure.ReadLogData("./../../game_score_log.csv")
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
			want:nil,
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