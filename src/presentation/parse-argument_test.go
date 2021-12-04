package presentation

import (
	"reflect"
	"testing"
)

func TestParseArgment(t *testing.T) {
	tests := []struct {
		name string
		want Argment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseArgment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgment() = %v, want %v", got, tt.want)
			}
		})
	}
}
