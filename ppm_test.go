package pnm

import "testing"

func TestShouldCreateNewPpmImage(t *testing.T) {
	type args struct {
		magicNumber int64
		maxValue    byte
		width       int64
		height      int64
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{"Should create new ppm image A", args{PPMText, 255, 8, 8}, "TestShouldCreateNewPpmImage.golden"},
		{"Should create new ppm image b", args{PPMText, 0, 0, 0}, "TestShouldCreateNewPpmImageb.golden"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewPPMImage(tc.args.width, tc.args.height, tc.args.maxValue, PPMText).String()
			want := OpenGoldenFile(t, tc.want, got, *update)
			if want != got {
				t.Errorf("NewPPMImage = %v want %v", got, want)
			}
		})
	}
}
