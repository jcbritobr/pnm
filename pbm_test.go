package pnm

import (
	"flag"
	"reflect"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestShouldCreateNewPbmImage(t *testing.T) {
	type args struct {
		magicNumber int64
		width       int64
		height      int64
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{"Should create new pbm image A", args{PBMText, 8, 8}, "TestShouldCreateNewPbmImage.golden"},
		{"Should create new pbm image b", args{PBMText, 0, 0}, "TestShouldCreateNewPbmImageb.golden"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewPBMImage(tc.args.width, tc.args.height, PBMText).String()
			want := OpenGoldenFile(t, tc.want, got, *update)
			if want != got {
				t.Errorf("NewPBMImage = %v want %v", got, want)
			}
		})
	}
}

func TestPbmImageGettersShouldReturnCorrectly(t *testing.T) {
	type args struct {
		mn     int64
		mv     byte
		width  int64
		height int64
	}
	testCases := []struct {
		name    string
		args    args
		wmn     int64
		wmv     byte
		wwidth  int64
		wheight int64
		mbuf    []byte
	}{
		{"Pbm image getter should return correctly A", args{PBMText, 1, 8, 8}, PBMText, 0, 8, 8, make([]byte, 64)},
		{"Pbm image getter should return correctly B", args{PBMText, 1, 0, 0}, PBMText, 0, 0, 0, make([]byte, 0)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			image := NewPBMImage(tc.args.width, tc.args.height, PBMText)

			if image.MagicNumber() != tc.wmn {
				t.Errorf("MagicNumber() = %v want %v", image.MagicNumber(), tc.wmn)
			}

			if image.Value() != tc.wmv {
				t.Errorf("Value() = %v want %v", image.Value(), tc.wmv)
			}

			if image.Width() != tc.wwidth {
				t.Errorf("Width() = %v want %v", image.Width(), tc.wwidth)
			}

			if image.Height() != tc.wheight {
				t.Errorf("Height() = %v want %v", image.Height(), tc.wheight)
			}

			if !reflect.DeepEqual(image.Buffer(), tc.mbuf) {
				t.Errorf("Buffer() = %v want %v", image.Buffer(), tc.mbuf)
			}
		})
	}
}
