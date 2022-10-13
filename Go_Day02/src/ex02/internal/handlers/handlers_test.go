package handlers

import "testing"

func TestHandler(t *testing.T) {
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Empty case",
			args: args{
				command: "",
				args:    nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "ls case1",
			args: args{
				command: "ls",
				args: []string{
					"-al",
					"a",
					"b",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "ls case2",
			args: args{
				command: "ls",
				args: []string{
					"../../internal",
					"../../cmd",
				},
			},
			want:    "../../cmd:\nmain.go\n\n../../internal:\nhandlers\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.args.command, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}
