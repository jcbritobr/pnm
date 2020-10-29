package buffer

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/jcbritobr/pnm"
)

func TestImageBuffer(t *testing.T) {
	t.Run("Should create NewImageBuffer", func(t *testing.T) {
		want := &ImageBuffer{offset: 0, pix: make([]byte, 100*100*3)}
		got := NewImageBuffer(pnm.NewPPMImage(100, 100, 255, pnm.PPMBinary))
		if !reflect.DeepEqual(*want, *got) {
			t.Errorf("NewImageBuffer() = %v want %v", got, want)
		}
	})

	t.Run("Should read method interface read data into byte buffer correctly", func(t *testing.T) {
		b := make([]byte, 5*5*3)
		buffer := bytes.NewBuffer(b)
		want := NewImageBuffer(pnm.NewPPMImage(5, 5, 255, pnm.PPMBinary))
		_, err := io.Copy(buffer, want)
		if err != nil && !reflect.DeepEqual(b, want.pix) {
			t.Errorf("Write() = %v want %v", buffer, want.pix)
		}
	})

	t.Run("Should write method interface write data into ImageBuffer correctly", func(t *testing.T) {
		b := make([]byte, 5*5*3)
		buffer := bytes.NewBuffer(b)
		want := NewImageBuffer(pnm.NewPPMImage(5, 5, 255, pnm.PPMBinary))
		_, err := io.Copy(want, buffer)
		if err != nil && !reflect.DeepEqual(b, want.pix) {
			t.Errorf("Write() = %v want %v", buffer, want.pix)
		}
	})

	t.Run("Should write method interface fail with eof", func(t *testing.T) {
		b := make([]byte, 0)
		buffer := bytes.NewBuffer(b)
		want := NewImageBuffer(pnm.NewPPMImage(5, 5, 255, pnm.PPMBinary))
		_, err := want.Write(buffer.Bytes())
		if err == nil {
			t.Errorf("Write() = %v want %v", buffer, want.pix)
		}
	})
}
