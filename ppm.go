package pnm

// PPMImage implements a portable gray bitmap. This strucute has
// a gray value, compared with his binary option, PBM format
type PPMImage struct {
	*portableAnyMapImage
}

// NewPPMImage creates a new PGM image
func NewPPMImage(w, h int64, mv byte, t int64) *PPMImage {
	image := &PPMImage{
		portableAnyMapImage: &portableAnyMapImage{
			width:       w,
			height:      h,
			maxValue:    mv,
			magicNumber: t,
			buffer:      make([]byte, w*h*3),
		},
	}
	return image
}
