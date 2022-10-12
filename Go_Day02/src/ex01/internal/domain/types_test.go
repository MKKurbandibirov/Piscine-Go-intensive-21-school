package domain

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		flags *Flags
	}
	tests := []struct {
		name    string
		args    args
		want    *Flags
		wantErr bool
	}{
		{
			name: "Empty Case",
			args: args{
				flags: &Flags{
					L: false,
					W: false,
					M: false,
				},
			},
			want: &Flags{
				L: false,
				W: true,
				M: false,
			},
			wantErr: false,
		},
		{
			name: "Empty Case",
			args: args{
				flags: &Flags{
					L: true,
					W: true,
					M: false,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Case",
			args: args{
				flags: &Flags{
					L: true,
					W: false,
					M: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Case",
			args: args{
				flags: &Flags{
					L: false,
					W: true,
					M: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Case",
			args: args{
				flags: &Flags{
					L: true,
					W: true,
					M: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Valid(tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Valid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
