package infrastructure

import (
	"testing"
)

func TestReadLogData(t *testing.T) {
	tests := []struct {
		name string
		args string
		want [][]string
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:"SuccessCheckReadLogFile",
			args: "./../../game_score_log.csv",
			wantErr: nil,
		},
		//{
		//	name:"FailCheckReadLogFile",
		//	args: "./../../game_score_log.csv",
		//	wantErr: errors.New(""),
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_,gotErr := ReadLogData(tt.args)
			if gotErr != tt.wantErr {
				//アサーション
				t.Errorf("ReadLogData is  error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_checkLogHeader(t *testing.T) {
	record,_:=ReadLogData("./../../game_score_log.csv")
	header, _ := record.Read()
	type args struct {
		header []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name:"checkLogFileHeader",
			args: args{header: header},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLogHeader(tt.args.header); got != tt.want {
				t.Errorf("checkLogHeader is error = %v",got)
			}
		})
	}
}