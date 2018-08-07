package hexe

// Color defines an 8-bit packed color value in the HSLuv color space.
type Color uint8

// HSL defines a color value in the HSLuv color space.
type HSL struct {
	H, S, L float64
}

// decoding constants
const (
	hueOff         = 9.0        // of [0, 16)
	satMin, satMax = 0.25, 0.85 // of [0, 1]
	lgtMin, lgtMax = 0.25, 0.85 // of [0, 1]
)

func (c Color) Decode() HSL {
	v := uint(c)

	// non-color shades
	if v < 16 {
		return HSL{0, 0, stepValue(v, 16, 0, 1)}
	}

	// saturated colors
	// binary layout: HHHH.CCLL
	v -= 16

	// get indexes
	hi := v >> 4
	si := (v >> 2) & 3
	li := v & 3

	// convert to values
	return HSL{
		stepValue(hi, 16, 0, 360) + hueOff,
		stepValue(si, 4, satMin, satMax),
		stepValue(li, 4, lgtMin, lgtMax),
	}
}
