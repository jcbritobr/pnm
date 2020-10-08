package pnm

// PGMImage implements a portable gray bitmap. This strucute has
// a gray value, compared with his binary option, PBM format
type PGMImage struct {
	*portableAnyMapImage
}

// NewPGMImage creates a new PGM image
func NewPGMImage(w, h int64, mv byte, t int64) *PGMImage {
	image := &PGMImage{
		portableAnyMapImage: newAnyMapImage(w, h, mv, t),
	}
	return image
}
