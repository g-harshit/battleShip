package stratergy

import (
	"testing"

	"github.com/battleShip/lib"
)

var validLeft = map[string]struct{}{
	lib.MakeCoordinateString(0, 0): {},
	lib.MakeCoordinateString(0, 1): {},
	lib.MakeCoordinateString(1, 0): {},
	lib.MakeCoordinateString(1, 1): {},
	lib.MakeCoordinateString(2, 0): {},
	lib.MakeCoordinateString(2, 1): {},
	lib.MakeCoordinateString(3, 0): {},
	lib.MakeCoordinateString(3, 1): {},
}

func TestValidate(t *testing.T) {
	type args struct {
		grid                   [][]string
		cordinates, validRange map[string]struct{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				grid: [][]string{{".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}},
				cordinates: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				},
				validRange: validLeft,
			},
			wantErr: false,
		},
		{
			name: "Error on invalid coordinates",
			args: args{
				grid: [][]string{{".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}},
				cordinates: map[string]struct{}{
					lib.MakeCoordinateString(-1, 0): {},
				},
				validRange: validLeft,
			},
			wantErr: true,
		},
		{
			name: "Error on invalid overlapping grid",
			args: args{
				grid: [][]string{{"SH", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}},
				cordinates: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				},
				validRange: validLeft,
			},
			wantErr: true,
		},
		{
			name: "Error on invalid grid size",
			args: args{
				grid: [][]string{{".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}},
				cordinates: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				},
				validRange: validLeft,
			},
			wantErr: true,
		},
		{
			name: "Error on invalid plot",
			args: args{
				grid: [][]string{{".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}},
				cordinates: map[string]struct{}{
					lib.MakeCoordinateString(2, 0): {},
				},
				validRange: validLeft,
			},
			wantErr: true,
		},
	}
	st := NewBasicStratergy(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := st.Validate(tt.args.grid, tt.args.cordinates, tt.args.validRange)
			if (err != nil) && !tt.wantErr {
				t.Errorf("Validate() error : %v, wantError: %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOverlapping(t *testing.T) {
	type args struct {
		cordinatesA, cordinatesB map[string]struct{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				cordinatesA: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				},
				cordinatesB: map[string]struct{}{
					lib.MakeCoordinateString(2, 0): {},
				},
			},
			wantErr: false,
		},
		{
			name: "Error on coordinates overlapping",
			args: args{
				cordinatesA: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				}, cordinatesB: map[string]struct{}{
					lib.MakeCoordinateString(0, 0): {},
				},
			},
			wantErr: true,
		},
		{
			name: "Error on invalid coordinate A",
			args: args{
				cordinatesB: map[string]struct{}{
					lib.MakeCoordinateString(2, 0): {},
				},
			},
			wantErr: true,
		},
	}
	st := NewBasicStratergy(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := st.ValidateOverlapping(tt.args.cordinatesA, tt.args.cordinatesB)
			if (err != nil) && !tt.wantErr {
				t.Errorf("Validate() error : %v, wantError: %v", err, tt.wantErr)
			}
		})
	}
}
