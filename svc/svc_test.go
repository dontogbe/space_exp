package svc_test

import (
	"testing"

	"github.com/dontogbe/space_exp/svc"
)

func Test_svc_GetLocation(t *testing.T) {
	type args struct {
		x   float64
		y   float64
		z   float64
		vel float64
	}
	tests := []struct {
		name         string
		args         args
		wantLocation float64
		wantErr      bool
	}{
		{
			name: "happy path",
			args: args{
				x:   123.12,
				y:   456.56,
				z:   789.89,
				vel: 20.0,
			},
			wantLocation: 1389.5700000000002,
			wantErr:      false,
		},
		{
			name: "negative velocity",
			args: args{
				x:   0,
				y:   0,
				z:   0,
				vel: -1.11,
			},
			wantLocation: 0,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := svc.NewSvc()
			gotLocation, err := s.GetLocation(tt.args.x, tt.args.y, tt.args.z, tt.args.vel)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLocation != tt.wantLocation {
				t.Errorf("GetLocation() gotLocation = %v, want %v", gotLocation, tt.wantLocation)
			}
		})
	}
}
