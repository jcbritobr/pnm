package pnm

import "fmt"

// Constants for image types
const (
	PBMText int64 = iota + 1
	PGMText
	PPMText
	PBMBinary
	PGMBinary
	PPMBinary
)

// Image represents all portable anytype format
type Image interface {
	MagicNumber() int64
	Width() int64
	Height() int64
	Buffer() []byte
	Value() byte
}

// PortableAnyMapImage implements a Portable anymap format. This
// struct will compose other.
type portableAnyMapImage struct {
	magicNumber int64
	width       int64
	height      int64
	maxValue    byte
	buffer      []byte
}

// NewPBMImage creates a new pbm image
func newAnyMapImage(w, h int64, mv byte, t int64) *portableAnyMapImage {
	image := &portableAnyMapImage{
		magicNumber: t,
		width:       w,
		height:      h,
		maxValue:    mv,
		buffer:      make([]byte, w*h),
	}
	return image
}

// MagicNumber return the magic number information of image
func (p *portableAnyMapImage) MagicNumber() int64 {
	return p.magicNumber
}

// Width returns the width of image
func (p *portableAnyMapImage) Width() int64 {
	return p.width
}

// Height returns the height of image
func (p *portableAnyMapImage) Height() int64 {
	return p.height
}

// Buffer returns the buffer of image
func (p *portableAnyMapImage) Buffer() []byte {
	return p.buffer
}

// Value returns the max value of image
func (p *portableAnyMapImage) Value() byte {
	return p.maxValue
}

// String implements interface Stringer
func (p *portableAnyMapImage) String() string {
	return fmt.Sprintf("mn:%v mv: %v width:%v height:%v buffer:%v", p.magicNumber, p.Value(), p.width, p.height, p.buffer)
}
