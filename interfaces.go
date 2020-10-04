package pnm

// Image represents all portable anytype format
type Image interface {
	MagicNumber() string
	Width() int
	Height() int
	Buffer() []byte
	Value() byte
}
