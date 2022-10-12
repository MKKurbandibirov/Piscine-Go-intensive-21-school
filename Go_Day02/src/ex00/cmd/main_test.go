package main

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	type args struct {
		fl Flags
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "All Case",
			args: args{fl: Flags{
				File:     false,
				Dir:      false,
				SymLink:  false,
				Ext:      "",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir",
				"../test_dir/sym_link",
				"../test_dir/test.txt",
				"../test_dir/test1.txt",
				"../test_dir/test_subdir",
				"../test_dir/test_subdir/test.go",
				"../test_dir/test_subdir/test_link",
			},
		},
		{
			name: "Only SymLink Case",
			args: args{fl: Flags{
				File:     false,
				Dir:      false,
				SymLink:  true,
				Ext:      "",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir/sym_link",
			},
		},
		{
			name: "Only Dir Case",
			args: args{fl: Flags{
				File:     false,
				Dir:      true,
				SymLink:  false,
				Ext:      "",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir",
				"../test_dir/test_subdir",
				"../test_dir/test_subdir/test_link",
			},
		},
		{
			name: "Only File Case",
			args: args{fl: Flags{
				File:     true,
				Dir:      false,
				SymLink:  false,
				Ext:      "",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir/test.txt",
				"../test_dir/test1.txt",
				"../test_dir/test_subdir/test.go",
			},
		},
		{
			name: "Ext Case",
			args: args{fl: Flags{
				File:     true,
				Dir:      false,
				SymLink:  false,
				Ext:      "txt",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir/test.txt",
				"../test_dir/test1.txt",
			},
		},
		{
			name: "Only Ext Case",
			args: args{fl: Flags{
				File:     false,
				Dir:      false,
				SymLink:  true,
				Ext:      "txt",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir/sym_link",
			},
		},
		{
			name: "Dir and File Case",
			args: args{fl: Flags{
				File:     true,
				Dir:      true,
				SymLink:  false,
				Ext:      "",
				RootName: "../test_dir",
			}},
			want: []string{
				"../test_dir",
				"../test_dir/test.txt",
				"../test_dir/test1.txt",
				"../test_dir/test_subdir",
				"../test_dir/test_subdir/test.go",
				"../test_dir/test_subdir/test_link",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exec(tt.args.fl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
