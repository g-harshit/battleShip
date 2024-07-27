package game

import (
	"testing"
)

func TestInitGame(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{2},
			wantErr: false,
		},
		{
			name:    "FAIL",
			args:    args{0},
			wantErr: true,
		},
		{
			name:    "FAIL",
			args:    args{1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InitGame(tt.args.n)
			if (err != nil) && !tt.wantErr {
				t.Errorf("InitGame() error : %v, wantError: %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddShip(t *testing.T) {
	type args struct {
		ID                   string
		size, xa, xb, ya, yb int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{"SH1", 2, 1, 1, 1, 4},
			wantErr: false,
		},
		{
			name:    "error out of bound",
			args:    args{"SH1", 2, 10, 12, 0, 4},
			wantErr: true,
		},
		{
			name:    "overlap",
			args:    args{"SH1", 2, 0, 0, 0, 4},
			wantErr: true,
		},
		{
			name:    "wrong side",
			args:    args{"SH1", 2, 4, 5, 0, 4},
			wantErr: true,
		},
		{
			name:    "wrong side",
			args:    args{"SH1", 2, 4, 0, 5, 5},
			wantErr: false,
		},
	}
	g, _ := InitGame(6)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := g.AddShip(tt.name, tt.args.size, tt.args.xa, tt.args.xb, tt.args.ya, tt.args.yb)
			if (err != nil) && !tt.wantErr {
				t.Errorf("InitGame() error : %v, wantError: %v", err, tt.wantErr)
			}
		})
	}
}
