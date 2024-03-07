package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v: got '%g' want '%g'", tt.shape, got, tt.want)
			}
		})
	}
}
