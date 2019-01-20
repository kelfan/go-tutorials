package logger

import "testing"

func TestInfo(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"info", args{"This is log information"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"warn", args{"This is warning information"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.msg)
		})
	}
}

func TestErr(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"error", args{"This is error information"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Err(tt.args.msg)
		})
	}
}
