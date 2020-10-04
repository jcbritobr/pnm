package pnm

// Image represents all portable anytype format
type Image interface {
	MagicNumber() string
	Width() int
	Height() int
	Buffer() []byte
	Value() byte
}

// PortableAnyMapImage implements a Portable anymap format. This
// struct will compose other.
type portableAnyMapImage struct {
	magicNumber string
	width       int
	height      int
	buffer      []byte
}
