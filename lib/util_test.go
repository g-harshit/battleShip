package lib

import "testing"

func TestGetCordinates(t *testing.T) {
	type args struct {
		coord string
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			name:  "success",
			args:  args{"1,3"},
			wantX: 1,
			wantY: 3,
		},
		{
			name:  "no commoa",
			args:  args{"13"},
			wantX: 0,
			wantY: 0,
		},
		{
			name:  "negative value",
			args:  args{"-1,13"},
			wantX: -1,
			wantY: 13,
		},
		{
			name:  "string value",
			args:  args{"hello,world"},
			wantX: 0,
			wantY: 0,
		},
		{
			name:  "empty value",
			args:  args{},
			wantX: 0,
			wantY: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := GetCordinates(tt.args.coord)
			if x != tt.wantX {
				t.Errorf("GetCordinates() X got : %v, want : %v", x, tt.wantX)
				return
			}
			if y != tt.wantY {
				t.Errorf("GetCordinates() Y got : %v, want : %v", y, tt.wantY)
				return
			}
		})
	}
}

func TestMakeCoordinateString(t *testing.T) {
	type args struct {
		x, y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{1, 2},
			want: "1,2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeCoordinateString(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("MakeCoordinateString() got : %v, want : %v", got, tt.want)
			}
		})
	}
}
