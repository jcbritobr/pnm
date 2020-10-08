package pnm

import "testing"

func TestShouldCreateNewPgmImage(t *testing.T) {
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
		{"Should create new pgm image A", args{PGMText, 255, 8, 8}, "TestShouldCreateNewPgmImage.golden"},
		{"Should create new pgm image b", args{PGMText, 0, 0, 0}, "TestShouldCreateNewPgmImageb.golden"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewPGMImage(tc.args.width, tc.args.height, tc.args.maxValue, PGMText).String()
			want := OpenGoldenFile(t, tc.want, got, *update)
			if want != got {
				t.Errorf("NewPGMImage = %v want %v", got, want)
			}
		})
	}
}
