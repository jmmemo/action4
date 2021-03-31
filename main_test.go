package main

import "testing"

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test1", want: "0.0.1"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := GetVersion(); got != test.want {
				t.Errorf("GetVersion() = %v,want %v", got, test.want)
			}
		})
	}
}
