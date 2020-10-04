package pnm

// PBMImage implements a Portable bitmap. Its a binary
// image, composed by 0 1 data
type PBMImage struct {
	portableAnyMapImage
}

// Width returns the width of image
func (p *PBMImage) Width() int {
	return p.width
}

// Height returns the height of image
func (p *PBMImage) Height() int {
	return p.height
}
