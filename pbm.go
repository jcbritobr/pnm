package pnm

// PBMImage implements a Portable Bitmap format. A binary image
// with values 0 or 1.
type PBMImage struct {
	magicNumber string
	width       int
	height      int
	buffer      []byte
}
