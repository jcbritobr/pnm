package integration_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/jcbritobr/pnm"
	"github.com/jcbritobr/pnm/buffer"
	"github.com/jcbritobr/pnm/color"
)

func TestPGMImage(t *testing.T) {
	t.Run("Should pgm image and encoder synthesizes a random image function", func(t *testing.T) {
		image := pnm.NewPGMImage(800, 800, 255, pnm.PGMBinary)
		file, err := os.Create("testdata/synimage.pgm")
		if err != nil {
			t.Errorf("fail with %v", err)
		}
		encoder := pnm.NewEncoder(file)

		imgbuf := image.Buffer()

		for i := range image.Buffer() {
			data := rand.Intn(255-0) + 0
			imgbuf[i] = byte(data)
		}

		err = encoder.Encode(image)
		if err != nil {
			t.Errorf("fail to encode image %v", err)
		}
	})

	t.Run("Should decode pgm generated image", func(t *testing.T) {
		var image pnm.PGMImage
		file, err := os.Open("testdata/synimage.pgm")
		if err != nil {
			t.Errorf("fail with %v", err)
		}
		decoder := pnm.NewDecoder(file, pnm.PGMBinary)
		err = decoder.Decode(&image)

		if err != nil {
			t.Errorf("fail to decode image %v", err)
		}
	})

	t.Run("Should decode and encode ppm image correctly", func(t *testing.T) {

		var image pnm.PPMImage
		file, err := os.Open("testdata/tree_1.ppm")
		defer file.Close()
		if err != nil {
			t.Errorf("fail with %v", err)
		}
		decoder := pnm.NewDecoder(file, pnm.PPMBinary)
		err = decoder.Decode(&image)

		if err != nil {
			t.Errorf("fail to decode image %v", err)
		}

		fe, err := os.Create("testdata/tree_2.ppm")
		if err != nil {
			t.Errorf("fail to create file %v", err)
		}
		defer fe.Close()
		encoder := pnm.NewEncoder(fe)
		err = encoder.Encode(&image)
		if err != nil {
			t.Errorf("Encode() = %v want %v", err, nil)
		}
	})

	t.Run("Should decode, process with ImageBuffer and RGBA color model", func(t *testing.T) {
		var image pnm.PPMImage
		file, err := os.Open("testdata/tree_1.ppm")
		if err != nil {
			t.Errorf("fail with %v", err)
		}
		defer file.Close()

		decoder := pnm.NewDecoder(file, pnm.PPMBinary)
		err = decoder.Decode(&image)
		if err != nil {
			t.Errorf("fail to decode image %v", err)
		}

		imbuf := buffer.NewImageBuffer(image)
		rgba := color.RGBA{}
		newbuf := []byte{}
		for {
			n, _ := imbuf.Read(rgba[:])
			if n <= 0 {
				break
			}
			r, g, b, _ := rgba.RGBA()
			r = 255 - r
			g = 255 - g
			b = 255 - b
			newbuf = append(newbuf, r)
			newbuf = append(newbuf, g)
			newbuf = append(newbuf, b)
		}

		image.SetBuffer(newbuf)

		fe, err := os.Create("testdata/tree_3.ppm")
		if err != nil {
			t.Errorf("fail to create file %v", err)
		}
		defer fe.Close()

		encoder := pnm.NewEncoder(fe)
		err = encoder.Encode(&image)
		if err != nil {
			t.Errorf("Encode() = %v want %v", err, nil)
		}
	})

}
