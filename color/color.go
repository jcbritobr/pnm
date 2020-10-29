package color

// Model represents the interface rgba color model
type Model interface {
	RGBA() (byte, byte, byte, byte)
}

// RGBA implements the RGBA color model interface
type RGBA [3]byte

// RGBA implements Model RGBA method
func (r RGBA) RGBA() (byte, byte, byte, byte) {
	return r[0], r[1], r[2], 0xff
}

// NewRGBA creates a new RGBA color model
func NewRGBA() RGBA {
	return RGBA{}
}

// Luma implements the Gray level color model interface
type Luma [1]byte

// RGBA implements Model RGBA method
func (l Luma) RGBA() (byte, byte, byte, byte) {
	return l[0], l[0], l[0], 0xff
}

// NewLuma creates a new Luma color model
func NewLuma() *Luma {
	return &Luma{}
}
