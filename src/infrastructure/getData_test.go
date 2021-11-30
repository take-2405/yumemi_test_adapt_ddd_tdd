package infrastructure

import (
	"testing"
)

func TestReadLogData(t *testing.T) {
	tests := []struct {
		name string
		args string
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:"checkReadLogFile",
			args: "./game_score_log.csv",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr!=nil {
				t.Errorf("ReadLogData is  error = %v, wantErr %v", nil, tt.wantErr)
				return
			}
		})
	}
}