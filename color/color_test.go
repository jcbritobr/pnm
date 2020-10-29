package color

import "testing"

func TestRGBA(t *testing.T) {
	t.Run("Should create a new rgb color model", func(t *testing.T) {
		rgba := NewRGBA()
		if len(rgba) != 3 {
			t.Errorf("NewRGB() = %v want allocated object with size 3", len(rgba))
		}
	})

	t.Run("Should create a new luma color model", func(t *testing.T) {
		luma := NewLuma()
		if len(luma) != 1 {
			t.Errorf("NewLuma() = %v want allocated object with size 1", len(luma))
		}
	})

	t.Run("Should read rgba values", func(t *testing.T) {
		rgba := NewRGBA()
		r, g, b, a := rgba.RGBA()
		if r != 0 || g != 0 || b != 0 || a != 0xff {
			t.Errorf("RGBA() = %v wants reseted buffer", rgba)
		}
	})

	t.Run("Should read luma values", func(t *testing.T) {
		luma := NewLuma()
		r, g, b, a := luma.RGBA()
		if r != 0 || g != 0 || b != 0 || a != 0xff {
			t.Errorf("RGBA() = %v wants reseted buffer", luma)
		}
	})
}
