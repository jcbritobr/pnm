package integration_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/jcbritobr/pnm"
)

func TestPGMImage(t *testing.T) {
	t.Run("Should pgm image and encoder synthesizes a random image function", func(t *testing.T) {
		image := pnm.NewPGMImage(800, 800, 255, pnm.PGMBinary)
		file, err := os.Create("testdata/sinimage.pgm")
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
}
