package testing

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	type args struct {
		f  float64
		f2 float64
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{"test", args{2.0, 2.0}, 8.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.f, tt.args.f2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
