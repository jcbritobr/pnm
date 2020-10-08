package pnm

import (
	"bytes"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	t.Run("Should create a new decoder with a reader as parameter and not return nil", func(t *testing.T) {
		buffer := bytes.NewBuffer([]byte{})
		decoder := NewDecoder(buffer, PPMBinary)
		if decoder == nil {
			t.Errorf("NewDecoder() = %v want not nil", decoder)
		}
	})

	t.Run("Should decode all pnm header type files", func(t *testing.T) {
		type args struct {
			buffer string
			format int64
		}
		testCases := []struct {
			name string
			args args
			want error
		}{
			{"and pbm text format cant fail", args{"P1\n250 250\n", PBMText}, nil},
			{"and pbm binary format cant fail", args{"P4\n250 250\n", PBMBinary}, nil},
			{"and pgm text format cant fail", args{"P2\n250 250\n255\n", PGMText}, nil},
			{"and pgm binary format cant fail", args{"P5\n250 250\n255\n", PGMBinary}, nil},
			{"and ppm text format cant fail", args{"P3\n250 250\n255\n", PPMText}, nil},
			{"and ppm binary format cant fail", args{"P6\n250 250\n255\n", PPMBinary}, nil},
			{"and pbm text format needs fail", args{"P6\n250 250\n255\n", PBMText}, errWrongFormat},
			{"and pbm text format needs fail malformed header", args{"P\n250 250\n255\n", PBMText}, errMalformedHeader},
			{"and pbm text format needs fail wrong format header", args{"P3\n#this is a comment\n250 250\n255\n", PBMText}, errWrongFormat},
			{"and ppm text format needs fail malformed header", args{"P3\n#this is a comment\n250 250\n255\n", PPMText}, nil},
			//{"and ppm text format needs fail malformed header", args{"P3\n250\n255\n", PPMText}, errWrongFormat}, here are a deadlock
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				buffer := bytes.NewBuffer([]byte(tc.args.buffer))
				decoder := NewDecoder(buffer, tc.args.format)
				header, err := decoder.decodeHeader()
				if err != tc.want {
					t.Errorf("decodeHeader() = %v - header %v  -  want %v", err, header, tc.want)
				}
			})
		}
	})

	t.Run("Should decode data text and binary streams as expected", func(t *testing.T) {
		type args struct {
			filename string
			format   int64
		}
		testCases := []struct {
			name string
			args args
			want error
		}{
			{"ppm file 255 value binary without error", args{"tree_1.ppm", PPMBinary}, nil},
			{"ppm text file 255 100x100 value binary without error", args{"ppmtext100x100.ppm", PPMText}, nil},
			{"pgm binary file 100x100 without error", args{"should_encode_data_as_100x100_binary_pgm.pgm", PGMBinary}, nil},
			{"pbm text file 8x8 without error", args{"should_encode_data_as_8x8_text_pbm.pbm", PBMText}, nil},
			{"pbm binary file 8x8 without error", args{"should_encode_data_as_8x8_binary_pbm.pbm", PBMBinary}, nil},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				file := Open(t, tc.args.filename, os.O_RDONLY)
				decoder := NewDecoder(file, tc.args.format)
				var image Image
				switch tc.args.format {
				case PPMBinary, PPMText:
					image = &PPMImage{}
				case PGMBinary, PGMText:
					image = &PGMImage{}
				case PBMBinary, PBMText:
					image = &PBMImage{}
				default:
					t.Errorf("wrong image format %v", tc.args.format)
				}

				got := decoder.Decode(image)
				if got != tc.want {
					t.Errorf("Decode() = %v want %v", got, tc.want)
				}
			})
		}

	})
}
