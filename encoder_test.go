package pnm

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewEncoderShouldCreateSucceed(t *testing.T) {
	got := NewEncoder(os.Stdout)
	if got == nil {
		t.Errorf("NewEncoder() = %v want %v", got, nil)
	}
}

func TestCheckBinaryShouldReturnAsExpected(t *testing.T) {
	type args struct {
		format int64
	}
	testCases := []struct {
		name string
		args args
		want bool
	}{
		{"CheckBinary should return true", args{PBMBinary}, true},
		{"CheckBinary should return false", args{PBMText}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encoder := NewEncoder(os.Stdout)
			got := encoder.checkBinaryFormat(tc.args.format)
			fmt.Println(got)
			if got != tc.want {
				t.Errorf("checkBinaryFormat() = %v want %v", got, tc.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		image Image
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{"should encode data as 8x8 binary pbm", args{NewPBMImage(8, 8, PBMBinary)}, "should_encode_data_as_8x8_binary_pbm.pbm"},
		{"should encode data as 8x8 text pbm", args{NewPBMImage(8, 8, PBMText)}, "should_encode_data_as_8x8_text_pbm.pbm"},
		{"should encode data as 100x100 binary pgm", args{NewPGMImage(100, 100, 15, PGMBinary)}, "should_encode_data_as_100x100_binary_pgm.pgm"},
		{"should encode data as 100x100 binary ppm", args{NewPPMImage(100, 100, 255, PPMBinary)}, "should_encode_data_as_100x100_binary_ppm.ppm"},
		{"should encode data as 100x100 text ppm", args{NewPPMImage(100, 100, 255, PPMText)}, "should_encode_data_as_100x100_text_ppm.ppm"},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := bytes.NewBuffer([]byte{})
			encoder := NewEncoder(got)
			err := encoder.Encode(tC.args.image)
			if err != nil {
				t.Errorf("Encoder() = %v", err)
			}

			file := Open(t, tC.want, os.O_RDWR)
			file.(*os.File).Truncate(0)

			_, err = io.Copy(file, got)
			if err != nil {
				t.Errorf("Encoder() = %v", err)
			}
		})
	}

	t.Run("Should encode pass with a ruge buffer", func(t *testing.T) {
		image := NewPPMImage(10000, 10000, 255, PPMBinary)
		buffer := bytes.NewBuffer([]byte{})
		encoder := NewEncoder(buffer)

		err := encoder.Encode(image)
		if err != nil {
			t.Errorf("Buffer size is too large")
		}
	})
}
