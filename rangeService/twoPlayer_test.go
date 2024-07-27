package rangeService

import (
	"reflect"
	"testing"

	"github.com/battleShip/lib"
)

func TestGetPlayerRange(t *testing.T) {
	type args struct {
		player string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "Success Player A",
			args: args{lib.PlayerA},
			want: map[string]struct{}{
				lib.MakeCoordinateString(0, 0): {},
				lib.MakeCoordinateString(0, 1): {},
				lib.MakeCoordinateString(1, 0): {},
				lib.MakeCoordinateString(1, 1): {},
				lib.MakeCoordinateString(2, 0): {},
				lib.MakeCoordinateString(2, 1): {},
				lib.MakeCoordinateString(3, 0): {},
				lib.MakeCoordinateString(3, 1): {},
			},
		},
		{
			name: "Success Player B",
			args: args{lib.PlayerB},
			want: map[string]struct{}{
				lib.MakeCoordinateString(0, 2): {},
				lib.MakeCoordinateString(0, 3): {},
				lib.MakeCoordinateString(1, 2): {},
				lib.MakeCoordinateString(1, 3): {},
				lib.MakeCoordinateString(2, 2): {},
				lib.MakeCoordinateString(2, 3): {},
				lib.MakeCoordinateString(3, 2): {},
				lib.MakeCoordinateString(3, 3): {},
			},
		},
	}
	rs := NewTwoPlayerRangeService(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rs.GetPlayerRange(tt.args.player)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayerRange : \ngot %v \nwant %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomCordinates(t *testing.T) {
	type args struct {
		player string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success Player A",
			args:    args{lib.PlayerA},
			wantErr: false,
		},
		{
			name:    "Success Player B",
			args:    args{lib.PlayerB},
			wantErr: false,
		},
		{
			name:    "Success Player A",
			args:    args{lib.PlayerA},
			wantErr: false,
		},
		{
			name:    "Error Player A",
			args:    args{lib.PlayerA},
			wantErr: true,
		},
	}
	rs := NewTwoPlayerRangeService(2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := rs.GetRandomCordinates(tt.args.player)
			if (err != nil) && !tt.wantErr {
				t.Errorf("GetRandomCordinates() error : %v, wantError: %v", err, tt.wantErr)
			}
		})
	}
}
