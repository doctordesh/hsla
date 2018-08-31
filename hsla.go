// The conversions are not completely 100% since there are
// rounding errors that occur during the conversions. This
// is something that probably can't be fixed since it seem
// to be an inherent problem between the RGB and HSL color
// representations.
package hsla

import (
    "image/color"
    "math"
)

type HSLA struct {
    H float64
    S float64
    L float64
    A float64
}

// Creates a new HSLA struct from the built in color.RGBA type
func NewHSLAFromRGBA(c color.RGBA) HSLA {
    hsla := HSLA{}

    r := float64(c.R)
    g := float64(c.G)
    b := float64(c.B)

    max := math.Max(math.Max(r, g), b)
    min := math.Min(math.Min(r, g), b)

    // Luminosity is the average of the max and min rgb color intensities.
    hsla.L = (max + min) / 2

    // saturation
    delta := max - min
    if delta == 0 {
        // it's gray
        return hsla
    }

    // it's not gray
    if hsla.L < 0.5 {
        hsla.S = delta / (max + min)
    } else {
        hsla.S = delta / (2 - max - min)
    }

    // hue
    r2 := (((max - r) / 6) + (delta / 2)) / delta
    g2 := (((max - g) / 6) + (delta / 2)) / delta
    b2 := (((max - b) / 6) + (delta / 2)) / delta
    switch {
    case r == max:
        hsla.H = b2 - g2
    case g == max:
        hsla.H = (1.0 / 3.0) + r2 - b2
    case b == max:
        hsla.H = (2.0 / 3.0) + g2 - r2
    }

    // fix wraparounds
    switch {
    case hsla.H < 0:
        hsla.H += 1
    case hsla.H > 1:
        hsla.H -= 1
    }

    return hsla
}

// Converts a HSLA struct to the built in color.RGBA type.
func (c HSLA) ToRGBA() color.RGBA {
    h := c.H
    s := c.S
    l := c.L

    if s == 0 {
        return rgbaFromFloats(l, l, l, c.A)
    }

    var v1, v2 float64
    if l < 0.5 {
        v2 = l * (1 + s)
    } else {
        v2 = (l + s) - (s * l)
    }

    v1 = 2*l - v2

    r := hueToRGB(v1, v2, h+(1.0/3.0))
    g := hueToRGB(v1, v2, h)
    b := hueToRGB(v1, v2, h-(1.0/3.0))

    return rgbaFromFloats(r, g, b, c.A)
}

func hueToRGB(v1, v2, h float64) float64 {
    if h < 0 {
        h += 1
    }
    if h > 1 {
        h -= 1
    }
    switch {
    case 6*h < 1:
        return (v1 + (v2-v1)*6*h)
    case 2*h < 1:
        return v2
    case 3*h < 2:
        return v1 + (v2-v1)*((2.0/3.0)-h)*6
    }
    return v1
}

func rgbaFromFloats(r, g, b, a float64) color.RGBA {
    c := color.RGBA{
        uint8(math.Round(r * 255)),
        uint8(math.Round(g * 255)),
        uint8(math.Round(b * 255)),
        uint8(math.Round(a * 255)),
    }

    return c
}
