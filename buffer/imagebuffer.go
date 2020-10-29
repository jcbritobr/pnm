package buffer

import (
	"io"

	"github.com/jcbritobr/pnm"
)

// ImageBuffer is a buffer for processing data
type ImageBuffer struct {
	pix    []byte
	offset uint
}

// NewImageBuffer creates a new image buffer
func NewImageBuffer(image pnm.Image) *ImageBuffer {
	return &ImageBuffer{pix: image.Buffer(), offset: 0}
}

// Empty checks is the buffer was consumed
func (i *ImageBuffer) Empty() bool {
	return int(i.offset) >= len(i.pix)
}

func (i *ImageBuffer) Read(p []byte) (int, error) {
	if i.Empty() {
		return 0, io.EOF
	}

	n := copy(p, i.pix[i.offset:])
	i.offset += uint(n)
	return n, nil
}

func (i *ImageBuffer) Write(p []byte) (int, error) {
	if len(p) <= 0 {
		return 0, io.EOF
	}

	n := copy(i.pix, p)

	return n, nil
}
