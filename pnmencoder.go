package pnm

import (
	"bytes"
	"fmt"
	"io"
)

// Encoder encodes a pnm type image and writes it to a device
// the implements Writer interface
type Encoder struct {
	writer io.Writer
}

// NewEncoder creates a new Encoder
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer}
}

// checkBinaryFormat checks if magic number is a binary format
func (e *Encoder) checkBinaryFormat(format int64) bool {
	switch format {
	case PPMBinary, PBMBinary, PGMBinary:
		return true
	default:
		return false
	}
}

// getHeader gets a formatted image header
func (e *Encoder) getHeader(mn, w, h int64, mv byte) string {
	if mv > 1 {
		return fmt.Sprintf("P%d\n%d %d\n%d\n", mn, w, h, mv)
	}
	return fmt.Sprintf("P%d\n%d %d\n", mn, w, h)
}

// Encode encodes image data to device
func (e *Encoder) Encode(image Image) error {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteString(e.getHeader(image.MagicNumber(), image.Width(), image.Height(), image.Value()))
	if e.checkBinaryFormat(image.MagicNumber()) {
		_, err := buffer.Write(image.Buffer())
		if err != nil {
			return err
		}
	} else {

		for _, b := range image.Buffer() {
			_, err := buffer.WriteString(fmt.Sprintf("%d", b))
			if err != nil {
				return err
			}
			_, err = buffer.WriteString(" ")
			if err != nil {
				return err
			}
		}
	}
	io.Copy(e.writer, buffer)
	return nil
}
