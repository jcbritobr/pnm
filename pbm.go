package pnm

import "fmt"

// PBMImage implements a Portable bitmap. Its a binary
// image, composed by 0 1 data
type PBMImage struct {
	*portableAnyMapImage
}

// NewPBMImage creates a new pbm image
func NewPBMImage(w, h, t int64) *PBMImage {
	image := &PBMImage{
		portableAnyMapImage: newAnyMapImage(w, h, 0, t),
	}
	return image
}

// String implements interface Stringer
func (p *PBMImage) String() string {
	return fmt.Sprintf("mn:%v mv: %v width:%v height:%v buffer:%v", p.magicNumber, p.Value(), p.width, p.height, p.buffer)
}
