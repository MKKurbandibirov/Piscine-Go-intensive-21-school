package handler

import "testing"

func TestHandleWords(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty Case",
			args: args{
				fileName: "../../test/empty.txt",
			},
			want: 1,
		},
		{
			name: "Russian Case",
			args: args{
				fileName: "../../test/input1.txt",
			},
			want: 48,
		},
		{
			name: "English Case",
			args: args{
				fileName: "../../test/input2.txt",
			},
			want: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleWords(tt.args.fileName); got != tt.want {
				t.Errorf("HandleWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleLines(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty Case",
			args: args{
				fileName: "../../test/empty.txt",
			},
			want: 1,
		},
		{
			name: "Russian Case",
			args: args{
				fileName: "../../test/input1.txt",
			},
			want: 4,
		},
		{
			name: "English Case",
			args: args{
				fileName: "../../test/input2.txt",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleLines(tt.args.fileName); got != tt.want {
				t.Errorf("HandleLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleSymbols(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty Case",
			args: args{
				fileName: "../../test/empty.txt",
			},
			want: 0,
		},
		{
			name: "Russian Case",
			args: args{
				fileName: "../../test/input1.txt",
			},
			want: 377,
		},
		{
			name: "English Case",
			args: args{
				fileName: "../../test/input2.txt",
			},
			want: 746,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleSymbols(tt.args.fileName); got != tt.want {
				t.Errorf("HandleSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
