package filer

import (
	"os"
	"strings"
	"testing"
)

func TestCheckFolder(t *testing.T) {
	type args struct {
		path           string
		permissionBits os.FileMode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{"./test1", 0700}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckFolder(tt.args.path, tt.args.permissionBits)
		})
	}
}

func TestCurrentPath(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"currentPath"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CurrentPath()
		})
	}
}

func TestRelative2abs(t *testing.T) {
	type args struct {
		rel string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Relative2Abs", args{"./"}, "util/filer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Relative2abs(tt.args.rel); !strings.Contains(got, tt.want) {
				t.Errorf("Relative2abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadfile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"readfile", args{"/usr/local/go/CONTRIBUTING.md"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Readfile(tt.args.file); got == tt.want {
				t.Errorf("Readfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
