package infrastructure

import (
	"encoding/csv"
	"log"
	"testing"
	"yumemi-coding-test-practice/src/domain"
	"yumemi-coding-test-practice/src/domain/repository"
)

func Test_checkLogHeader(t *testing.T) {
	//渡されたファイルパスのファイルを開く
	var successHeader []string = []string{"create_timestamp", "player_id", "score"}
	var faildHeader []string = []string{"a", "b", "c"}
	type args struct {
		header []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successCheckHeader",
			args: args{header: successHeader},
			want: true,
		},
		{
			name: "successCheckHeader",
			args: args{header: faildHeader},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLogHeader(tt.args.header); got != tt.want {
				t.Errorf("checkLogHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playDataPersistence_ParsePlayData(t *testing.T) {
	var PlayData domain.PlayDatas
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

	type fields struct {
		PlayData domain.PlayDatas
	}
	type args struct {
		csv *csv.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.PlayDatas
		wantErr error
	}{
		{
			name:    "successParsePlayData",
			fields:  fields{PlayData: PlayData},
			args:    args{csv: csv},
			want:    PlayData,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := playDataPersistence{
				PlayData: tt.fields.PlayData,
			}
			got, err := p.ParsePlayData(tt.args.csv)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ParsePlayData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Data) == 0 {
				t.Errorf("ParsePlayData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playDataPersistence_ReadPlayData(t *testing.T) {
	var PlayData domain.PlayDatas
	type fields struct {
		PlayData domain.PlayDatas
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *csv.Reader
		wantErr error
	}{
		{
			name:    "successReadData",
			fields:  fields{PlayData: PlayData},
			args:    args{filePath: "./../../game_score_log.csv"},
			want:    nil,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := playDataPersistence{
				PlayData: tt.fields.PlayData,
			}
			_, err := p.ReadPlayData(tt.args.filePath)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ReadPlayData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
