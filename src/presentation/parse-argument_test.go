package presentation

import (
	"testing"
)

func TestParseArgment(t *testing.T) {
	tests := []struct {
		name string
		want Argment
	}{
		// TODO: Add test cases.
		{
			name: "",
			want: Argment{FilePath: "game_score_log.csv "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseArgment(); got.FilePath == "" {
				t.Errorf("argment is null ")
			}
		})
	}
}
